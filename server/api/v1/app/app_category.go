package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	xxxReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AppCategoryApi struct {
}

var appCategoryService = service.ServiceGroupApp.AppServiceGroup.AppCategoryService

// CreateAppCategory 创建AppCategory
// @Tags AppCategory
// @Summary 创建AppCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppCategory true "创建AppCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appCategory/createAppCategory [post]
func (appCategoryApi *AppCategoryApi) CreateAppCategory(c *gin.Context) {
	var appCategory app.AppCategory
	_ = c.ShouldBindJSON(&appCategory)
	if err := appCategoryService.CreateAppCategory(appCategory); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppCategory 删除AppCategory
// @Tags AppCategory
// @Summary 删除AppCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppCategory true "删除AppCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appCategory/deleteAppCategory [delete]
func (appCategoryApi *AppCategoryApi) DeleteAppCategory(c *gin.Context) {
	var appCategory app.AppCategory
	_ = c.ShouldBindJSON(&appCategory)
	if err := appCategoryService.DeleteAppCategory(appCategory); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppCategoryByIds 批量删除AppCategory
// @Tags AppCategory
// @Summary 批量删除AppCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appCategory/deleteAppCategoryByIds [delete]
func (appCategoryApi *AppCategoryApi) DeleteAppCategoryByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := appCategoryService.DeleteAppCategoryByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppCategory 更新AppCategory
// @Tags AppCategory
// @Summary 更新AppCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppCategory true "更新AppCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appCategory/updateAppCategory [put]
func (appCategoryApi *AppCategoryApi) UpdateAppCategory(c *gin.Context) {
	var appCategory app.AppCategory
	_ = c.ShouldBindJSON(&appCategory)
	if err := appCategoryService.UpdateAppCategory(appCategory); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppCategory 用id查询AppCategory
// @Tags AppCategory
// @Summary 用id查询AppCategory
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppCategory true "用id查询AppCategory"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appCategory/findAppCategory [get]
func (appCategoryApi *AppCategoryApi) FindAppCategory(c *gin.Context) {
	var appCategory app.AppCategory
	_ = c.ShouldBindQuery(&appCategory)
	if err, reappCategory := appCategoryService.GetAppCategory(appCategory.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappCategory": reappCategory}, c)
	}
}

// GetAppCategoryList 分页获取AppCategory列表
// @Tags AppCategory
// @Summary 分页获取AppCategory列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query xxxReq.AppCategorySearch true "分页获取AppCategory列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appCategory/getAppCategoryList [get]
func (appCategoryApi *AppCategoryApi) GetAppCategoryList(c *gin.Context) {
	var pageInfo xxxReq.AppCategorySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := appCategoryService.GetAppCategoryInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
