package member

import (
	"github.com/gin-gonic/gin"
	"github.com/tongxz/xs-admin-vue/api/v1"
	"github.com/tongxz/xs-admin-vue/middleware"
)

type RechargeRouter struct {
}

// InitRechargeRouter 初始化 Recharge 路由信息
func (s *RechargeRouter) InitRechargeRouter(Router *gin.RouterGroup) {
	rechargeRouter := Router.Group("recharge").Use(middleware.OperationRecord())
	rechargeRouterWithoutRecord := Router.Group("recharge")
	var rechargeApi = v1.ApiGroupApp.MemberApiGroup.RechargeApi
	{
		rechargeRouter.POST("createRecharge", rechargeApi.CreateRecharge)             // 新建Recharge
		rechargeRouter.DELETE("deleteRecharge", rechargeApi.DeleteRecharge)           // 删除Recharge
		rechargeRouter.DELETE("deleteRechargeByIds", rechargeApi.DeleteRechargeByIds) // 批量删除Recharge
		rechargeRouter.PUT("updateRecharge", rechargeApi.UpdateRecharge)              // 更新Recharge
	}
	{
		rechargeRouterWithoutRecord.GET("findRecharge", rechargeApi.FindRecharge)       // 根据ID获取Recharge
		rechargeRouterWithoutRecord.GET("getRechargeList", rechargeApi.GetRechargeList) // 获取Recharge列表
	}
}
