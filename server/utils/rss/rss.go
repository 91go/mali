package rss

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/utils/helper/slice"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/redis"

	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"

	"github.com/gorilla/feeds"
)

// Feed 通用Feed
type Feed struct {
	URL, Author string
	UpdatedTime time.Time
	Title
}

type Title struct {
	Prefix string
	Name   string
}

type Item struct {
	UpdatedTime time.Time
	Enclosure   *feeds.Enclosure
	URL         string
	Title       string
	Contents    string
	ID          string
	Author      string
	Description string
}

// Rss 输出rss
func Rss(fe *Feed, items []*Item) string {
	if len(items) == 0 {
		feed := feeds.Feed{
			Title:   feedTitle(fe.Title),
			Link:    &feeds.Link{Href: fe.URL},
			Author:  &feeds.Author{Name: fe.Author},
			Updated: fe.UpdatedTime,
		}
		atom, _ := feed.ToRss()
		return atom
	}

	if fe.UpdatedTime.IsZero() {
		return feedWithoutTime(fe, items)
	}

	return rss(fe, items)
}

func rss(fe *Feed, items []*Item) string {
	feed := feeds.Feed{
		Title:   feedTitle(fe.Title),
		Link:    &feeds.Link{Href: fe.URL},
		Author:  &feeds.Author{Name: fe.Author},
		Updated: fe.UpdatedTime,
	}

	for key := range items {
		feed.Add(&feeds.Item{
			Title:       items[key].Title,
			Link:        &feeds.Link{Href: items[key].URL},
			Content:     items[key].Contents,
			Description: items[key].Description,
			Author:      &feeds.Author{Name: items[key].Author},
			Id:          items[key].ID,
			Enclosure:   items[key].Enclosure,
			Updated:     items[key].UpdatedTime,
		})
	}

	// 输出atom，跟rsshub保持一致
	atom, err := feed.ToAtom()
	if err != nil {
		global.GVA_LOG.Error("function rss() failed, rss generate failed", zap.String("url", fe.URL), zap.String("err", err.Error()))
		return ""
	}
	return atom
}

func feedTitle(tt Title) string {
	if tt.Name == "" {
		return tt.Prefix
	}
	return fmt.Sprintf("%s - %s", tt.Prefix, tt.Name)
}

// 处理没有提供更新时间的feed
// 根据item的UpdatedTime判断
// TODO items[i] = item
func feedWithoutTime(feed *Feed, items []*Item) string {
	clt := redis.NewClient(redis.Conn())

	m := []string{}
	for key := range items {
		m = append(m, items[key].URL, gtime.Now().TimestampStr())
	}
	// 判断key是否存在，不存在则直接set并返回
	if clt.Conn.Exists(redis.Ctx, feed.URL).Val() != 1 {
		err := clt.Conn.HSet(redis.Ctx, feed.URL, m).Err()
		if err != nil {
			fmt.Println(err)
			return ""
		}
		for key := range items {
			items[key].UpdatedTime = gtime.Now().Time
		}
		return rss(feed, items)
	}

	// 如果更新了，把新数据append进去，再返回
	res := checkIsUpdate(clt, feed, items)
	if len(res) != 0 {
		n := []string{}
		for _, re := range res {
			n = append(n, re, gtime.Now().TimestampStr())
		}
		clt.Conn.HSet(redis.Ctx, feed.URL, n)
	}

	// 获取更新item
	old := clt.Conn.HGetAll(redis.Ctx, feed.URL).Val()
	for key := range items {
		if search, ok := old[items[key].URL]; ok {
			items[key].UpdatedTime = gtime.NewFromTimeStamp(gconv.Int64(search)).Time
		} else {
			fmt.Println(items[key].URL, "key not exist")
		}
	}
	return rss(feed, items)
}

func checkIsUpdate(clt *redis.Client, feed *Feed, items []*Item) []string {
	// 通过对比相同name下的key，检查item是否更新
	old := clt.Conn.HKeys(redis.Ctx, feed.URL).Val()

	neo := []string{}
	for key := range items {
		neo = append(neo, items[key].URL)
	}
	return slice.DiffSlices(old, neo)
}

func GenDateID(tag, link string) string {
	return fmt.Sprintf("tag:%s,%s:%s", tag, gtime.Date(), link)
}

// 用来处理不会更新的feed
func GenFixedID(tag, link string) string {
	return fmt.Sprintf("tag:%s,%s", tag, link)
}
