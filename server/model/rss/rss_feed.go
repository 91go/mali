// 自动生成模板RssFeed
package rss

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// RssFeed 结构体
// 如果含有time.Time 请自行import time包
type RssFeed struct {
      global.GVA_MODEL
      CateId  *int `json:"cateId" form:"cateId" gorm:"column:cate_id;comment:rss分类;size:10;"`
      Url  string `json:"url" form:"url" gorm:"column:url;comment:rss的url;size:255;"`
      Keywords  string `json:"keywords" form:"keywords" gorm:"column:keywords;comment:关键字过滤，逗号连接;size:255;"`
      IsPause  *bool `json:"isPause" form:"isPause" gorm:"column:is_pause;comment:是否暂停（0不暂停1暂停）;"`
}


// TableName RssFeed 表名
func (RssFeed) TableName() string {
  return "rss_feed"
}
