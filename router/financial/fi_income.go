package financial

import (
	"github.com/gin-gonic/gin"
	"github.com/tongxz/xs-admin-vue/server/api/v1"
	"github.com/tongxz/xs-admin-vue/server/middleware"
)

type IncomeRouter struct {
}

// InitIncomeRouter 初始化 Income 路由信息
func (s *IncomeRouter) InitIncomeRouter(Router *gin.RouterGroup) {
	incomeRouter := Router.Group("income").Use(middleware.OperationRecord())
	incomeRouterWithoutRecord := Router.Group("income")
	var incomeApi = v1.ApiGroupApp.FinancialApiGroup.IncomeApi
	{
		incomeRouter.POST("createIncome", incomeApi.CreateIncome)             // 新建Income
		incomeRouter.DELETE("deleteIncome", incomeApi.DeleteIncome)           // 删除Income
		incomeRouter.DELETE("deleteIncomeByIds", incomeApi.DeleteIncomeByIds) // 批量删除Income
		incomeRouter.PUT("updateIncome", incomeApi.UpdateIncome)              // 更新Income
	}
	{
		incomeRouterWithoutRecord.GET("findIncome", incomeApi.FindIncome)       // 根据ID获取Income
		incomeRouterWithoutRecord.GET("getIncomeList", incomeApi.GetIncomeList) // 获取Income列表
	}
}
