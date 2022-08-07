package financial

import (
	"github.com/gin-gonic/gin"
	"github.com/tongxz/xs-admin-vue/api/v1"
	"github.com/tongxz/xs-admin-vue/middleware"
)

type IncomeRouter struct {
}

// InitIncomeRouter 初始化 Income 路由信息
func (s *IncomeRouter) InitIncomeRouter(Router *gin.RouterGroup) {
	incomeRouter := Router.Group("income").Use(middleware.OperationRecord())
	incomeRouterWithoutRecord := Router.Group("income")
	var incomeApi = v1.ApiGroupApp.FinancialApiGroup.IncomeApi
	{
		incomeRouter.POST("createIncome", incomeApi.CreateIncome)           // 新建Income
		incomeRouter.POST("deleteIncome", incomeApi.DeleteIncome)           // 删除Income
		incomeRouter.POST("deleteIncomeByIds", incomeApi.DeleteIncomeByIds) // 批量删除Income
		incomeRouter.PUT("updateIncome", incomeApi.UpdateIncome)            // 更新Income
	}
	{
		incomeRouterWithoutRecord.GET("findIncome", incomeApi.FindIncome)              // 根据ID获取Income
		incomeRouterWithoutRecord.GET("getIncomeList", incomeApi.GetIncomeList)        // 获取Income列表
		incomeRouterWithoutRecord.POST("getIncomeExcel", incomeApi.GetIncomeListExcel) // 获取Income列表
	}
}
