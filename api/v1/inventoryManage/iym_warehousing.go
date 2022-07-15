package inventoryManage

import (
	"github.com/gin-gonic/gin"
	"github.com/tongxz/xs-admin-vue/global"
	"github.com/tongxz/xs-admin-vue/model/common/request"
	"github.com/tongxz/xs-admin-vue/model/common/response"
	"github.com/tongxz/xs-admin-vue/model/inventoryManage"
	inventoryManageReq "github.com/tongxz/xs-admin-vue/model/inventoryManage/request"
	"github.com/tongxz/xs-admin-vue/service"
	"go.uber.org/zap"
)

type WarehousingApi struct {
}

var warehousingService = service.ServiceGroupApp.InventorymanageServiceGroup.WarehousingService

// CreateWarehousing 创建Warehousing
// @Tags Warehousing
// @Summary 创建Warehousing
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body inventoryManage.Warehousing true "创建Warehousing"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /warehousing/createWarehousing [post]
func (warehousingApi *WarehousingApi) CreateWarehousing(c *gin.Context) {
	var warehousing inventoryManage.Warehousing
	_ = c.ShouldBindJSON(&warehousing)
	if err := warehousingService.CreateWarehousing(warehousing); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWarehousing 删除Warehousing
// @Tags Warehousing
// @Summary 删除Warehousing
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body inventoryManage.Warehousing true "删除Warehousing"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /warehousing/deleteWarehousing [delete]
func (warehousingApi *WarehousingApi) DeleteWarehousing(c *gin.Context) {
	var warehousing inventoryManage.Warehousing
	_ = c.ShouldBindJSON(&warehousing)
	if err := warehousingService.DeleteWarehousing(warehousing); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWarehousingByIds 批量删除Warehousing
// @Tags Warehousing
// @Summary 批量删除Warehousing
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Warehousing"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /warehousing/deleteWarehousingByIds [delete]
func (warehousingApi *WarehousingApi) DeleteWarehousingByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := warehousingService.DeleteWarehousingByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWarehousing 更新Warehousing
// @Tags Warehousing
// @Summary 更新Warehousing
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body inventoryManage.Warehousing true "更新Warehousing"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /warehousing/updateWarehousing [put]
func (warehousingApi *WarehousingApi) UpdateWarehousing(c *gin.Context) {
	var warehousing inventoryManage.Warehousing
	_ = c.ShouldBindJSON(&warehousing)
	if err := warehousingService.UpdateWarehousing(warehousing); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWarehousing 用id查询Warehousing
// @Tags Warehousing
// @Summary 用id查询Warehousing
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query inventoryManage.Warehousing true "用id查询Warehousing"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /warehousing/findWarehousing [get]
func (warehousingApi *WarehousingApi) FindWarehousing(c *gin.Context) {
	var warehousing inventoryManage.Warehousing
	_ = c.ShouldBindQuery(&warehousing)
	if rewarehousing, err := warehousingService.GetWarehousing(warehousing.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewarehousing": rewarehousing}, c)
	}
}

// GetWarehousingList 分页获取Warehousing列表
// @Tags Warehousing
// @Summary 分页获取Warehousing列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query inventoryManageReq.WarehousingSearch true "分页获取Warehousing列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /warehousing/getWarehousingList [get]
func (warehousingApi *WarehousingApi) GetWarehousingList(c *gin.Context) {
	var pageInfo inventoryManageReq.WarehousingSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := warehousingService.GetWarehousingInfoList(pageInfo); err != nil {
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
} // GetWarehousingList 分页获取Warehousing列表
// @Tags Warehousing
// @Summary 分页获取Warehousing列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query inventoryManageReq.WarehousingSearch true "分页获取Warehousing列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /warehousing/getWarehousingList [get]
func (warehousingApi *WarehousingApi) GetWarehousingName(c *gin.Context) {
	var warehousing inventoryManage.Warehousing
	_ = c.ShouldBindQuery(&warehousing)
	if rewarehousing, err := warehousingService.GetWarehousingName(); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		global.GVA_LOG.Info("查询成功!", zap.Any("warehousing", rewarehousing))
		response.OkWithData(gin.H{"rewarehousing": rewarehousing}, c)
	}
}
