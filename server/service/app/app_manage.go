package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	xxxReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppManageService struct {
}

// CreateAppManage 创建AppManage记录
// Author [piexlmax](https://github.com/piexlmax)
func (appManageService *AppManageService) CreateAppManage(appManage app.AppManage) (err error) {
	err = global.GVA_DB.Create(&appManage).Error
	return err
}

// DeleteAppManage 删除AppManage记录
// Author [piexlmax](https://github.com/piexlmax)
func (appManageService *AppManageService) DeleteAppManage(appManage app.AppManage) (err error) {
	err = global.GVA_DB.Delete(&appManage).Error
	return err
}

// DeleteAppManageByIds 批量删除AppManage记录
// Author [piexlmax](https://github.com/piexlmax)
func (appManageService *AppManageService) DeleteAppManageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppManage{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppManage 更新AppManage记录
// Author [piexlmax](https://github.com/piexlmax)
func (appManageService *AppManageService) UpdateAppManage(appManage app.AppManage) (err error) {
	err = global.GVA_DB.Save(&appManage).Error
	return err
}

// GetAppManage 根据id获取AppManage记录
// Author [piexlmax](https://github.com/piexlmax)
func (appManageService *AppManageService) GetAppManage(id uint) (err error, appManage app.AppManage) {
	err = global.GVA_DB.Where("id = ?", id).First(&appManage).Error
	return
}

// GetAppManageInfoList 分页获取AppManage记录
// Author [piexlmax](https://github.com/piexlmax)
func (appManageService *AppManageService) GetAppManageInfoList(info xxxReq.AppManageSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppManage{})
	var appManages []app.AppManage
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appManages).Error
	return err, appManages, total
}
