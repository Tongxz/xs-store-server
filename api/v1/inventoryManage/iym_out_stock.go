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

type OutStockApi struct {
}

var stockService = service.ServiceGroupApp.InventorymanageServiceGroup.OutStockService

// CreateOutStock 创建OutStock
// @Tags OutStock
// @Summary 创建OutStock
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body inventoryManage.OutStock true "创建OutStock"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stock/createOutStock [post]
func (stockApi *OutStockApi) CreateOutStock(c *gin.Context) {
	var stock inventoryManage.OutStock
	_ = c.ShouldBindJSON(&stock)
	if err := stockService.CreateOutStock(stock); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOutStock 删除OutStock
// @Tags OutStock
// @Summary 删除OutStock
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body inventoryManage.OutStock true "删除OutStock"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stock/deleteOutStock [delete]
func (stockApi *OutStockApi) DeleteOutStock(c *gin.Context) {
	var stock inventoryManage.OutStock
	_ = c.ShouldBindJSON(&stock)
	if err := stockService.DeleteOutStock(stock); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOutStockByIds 批量删除OutStock
// @Tags OutStock
// @Summary 批量删除OutStock
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除OutStock"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /stock/deleteOutStockByIds [delete]
func (stockApi *OutStockApi) DeleteOutStockByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := stockService.DeleteOutStockByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOutStock 更新OutStock
// @Tags OutStock
// @Summary 更新OutStock
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body inventoryManage.OutStock true "更新OutStock"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /stock/updateOutStock [put]
func (stockApi *OutStockApi) UpdateOutStock(c *gin.Context) {
	var stock inventoryManage.OutStock
	_ = c.ShouldBindJSON(&stock)
	if err := stockService.UpdateOutStock(stock); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOutStock 用id查询OutStock
// @Tags OutStock
// @Summary 用id查询OutStock
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query inventoryManage.OutStock true "用id查询OutStock"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /stock/findOutStock [get]
func (stockApi *OutStockApi) FindOutStock(c *gin.Context) {
	var stock inventoryManage.OutStock
	_ = c.ShouldBindQuery(&stock)
	if restock, err := stockService.GetOutStock(stock.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restock": restock}, c)
	}
}

// GetOutStockList 分页获取OutStock列表
// @Tags OutStock
// @Summary 分页获取OutStock列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query inventoryManageReq.OutStockSearch true "分页获取OutStock列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stock/getOutStockList [get]
func (stockApi *OutStockApi) GetOutStockList(c *gin.Context) {
	var pageInfo inventoryManageReq.OutStockSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := stockService.GetOutStockInfoList(pageInfo); err != nil {
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
