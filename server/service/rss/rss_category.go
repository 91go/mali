package rss

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/rss"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    rssReq "github.com/flipped-aurora/gin-vue-admin/server/model/rss/request"
)

type RssCategoryService struct {
}

// CreateRssCategory 创建RssCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssCategoryService *RssCategoryService) CreateRssCategory(rssCategory rss.RssCategory) (err error) {
	err = global.GVA_DB.Create(&rssCategory).Error
	return err
}

// DeleteRssCategory 删除RssCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssCategoryService *RssCategoryService)DeleteRssCategory(rssCategory rss.RssCategory) (err error) {
	err = global.GVA_DB.Delete(&rssCategory).Error
	return err
}

// DeleteRssCategoryByIds 批量删除RssCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssCategoryService *RssCategoryService)DeleteRssCategoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]rss.RssCategory{},"id in ?",ids.Ids).Error
	return err
}

// UpdateRssCategory 更新RssCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssCategoryService *RssCategoryService)UpdateRssCategory(rssCategory rss.RssCategory) (err error) {
	err = global.GVA_DB.Save(&rssCategory).Error
	return err
}

// GetRssCategory 根据id获取RssCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssCategoryService *RssCategoryService)GetRssCategory(id uint) (err error, rssCategory rss.RssCategory) {
	err = global.GVA_DB.Where("id = ?", id).First(&rssCategory).Error
	return
}

// GetRssCategoryInfoList 分页获取RssCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (rssCategoryService *RssCategoryService)GetRssCategoryInfoList(info rssReq.RssCategorySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&rss.RssCategory{})
    var rssCategorys []rss.RssCategory
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&rssCategorys).Error
	return err, rssCategorys, total
}