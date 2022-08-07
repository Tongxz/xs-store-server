package member

import (
	"github.com/gin-gonic/gin"
	"github.com/tongxz/xs-admin-vue/api/v1"
	"github.com/tongxz/xs-admin-vue/middleware"
)

type MemberRouter struct {
}

// InitMemberRouter 初始化 Member 路由信息
func (s *MemberRouter) InitMemberRouter(Router *gin.RouterGroup) {
	memberRouter := Router.Group("member").Use(middleware.OperationRecord())
	memberRouterWithoutRecord := Router.Group("member")
	var memberApi = v1.ApiGroupApp.MemberApiGroup.MemberApi
	{
		memberRouter.POST("createMember", memberApi.CreateMember)           // 新建Member
		memberRouter.POST("deleteMember", memberApi.DeleteMember)           // 删除Member
		memberRouter.POST("deleteMemberByIds", memberApi.DeleteMemberByIds) // 批量删除Member
		memberRouter.PUT("updateMember", memberApi.UpdateMember)            // 更新Member
	}
	{
		memberRouterWithoutRecord.GET("findMember", memberApi.FindMember)       // 根据ID获取Member
		memberRouterWithoutRecord.GET("getMemberList", memberApi.GetMemberList) // 获取Member列表
		memberRouterWithoutRecord.GET("getMemberName", memberApi.GetMemberName) // 获取Member列表
	}
}
