package inventoryManage

import (
	"github.com/gin-gonic/gin"
	"github.com/tongxz/xs-admin-vue/api/v1"
	"github.com/tongxz/xs-admin-vue/middleware"
)

type OutStockRouter struct {
}

// InitOutStockRouter 初始化 OutStock 路由信息
func (s *OutStockRouter) InitOutStockRouter(Router *gin.RouterGroup) {
	stockRouter := Router.Group("stock").Use(middleware.OperationRecord())
	stockRouterWithoutRecord := Router.Group("stock")
	var stockApi = v1.ApiGroupApp.InventorymanageApiGroup.OutStockApi
	{
		stockRouter.POST("createOutStock", stockApi.CreateOutStock)             // 新建OutStock
		stockRouter.DELETE("deleteOutStock", stockApi.DeleteOutStock)           // 删除OutStock
		stockRouter.DELETE("deleteOutStockByIds", stockApi.DeleteOutStockByIds) // 批量删除OutStock
		stockRouter.PUT("updateOutStock", stockApi.UpdateOutStock)              // 更新OutStock
	}
	{
		stockRouterWithoutRecord.GET("findOutStock", stockApi.FindOutStock)       // 根据ID获取OutStock
		stockRouterWithoutRecord.GET("getOutStockList", stockApi.GetOutStockList) // 获取OutStock列表
	}
}
