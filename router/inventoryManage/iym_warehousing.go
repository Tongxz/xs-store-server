package inventoryManage

import (
	"github.com/gin-gonic/gin"
	"github.com/tongxz/xs-admin-vue/server/api/v1"
	"github.com/tongxz/xs-admin-vue/server/middleware"
)

type WarehousingRouter struct {
}

// InitWarehousingRouter 初始化 Warehousing 路由信息
func (s *WarehousingRouter) InitWarehousingRouter(Router *gin.RouterGroup) {
	warehousingRouter := Router.Group("warehousing").Use(middleware.OperationRecord())
	warehousingRouterWithoutRecord := Router.Group("warehousing")
	var warehousingApi = v1.ApiGroupApp.InventorymanageApiGroup.WarehousingApi
	{
		warehousingRouter.POST("createWarehousing", warehousingApi.CreateWarehousing)             // 新建Warehousing
		warehousingRouter.DELETE("deleteWarehousing", warehousingApi.DeleteWarehousing)           // 删除Warehousing
		warehousingRouter.DELETE("deleteWarehousingByIds", warehousingApi.DeleteWarehousingByIds) // 批量删除Warehousing
		warehousingRouter.PUT("updateWarehousing", warehousingApi.UpdateWarehousing)              // 更新Warehousing
	}
	{
		warehousingRouterWithoutRecord.GET("findWarehousing", warehousingApi.FindWarehousing)       // 根据ID获取Warehousing
		warehousingRouterWithoutRecord.GET("getWarehousingList", warehousingApi.GetWarehousingList) // 获取Warehousing列表
	}
}
