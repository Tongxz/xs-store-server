package financial

import (
	"github.com/gin-gonic/gin"
	"github.com/tongxz/xs-admin-vue/api/v1"
	"github.com/tongxz/xs-admin-vue/middleware"
)

type ExpensesRouter struct {
}

// InitExpensesRouter 初始化 Expenses 路由信息
func (s *ExpensesRouter) InitExpensesRouter(Router *gin.RouterGroup) {
	expensesRouter := Router.Group("expenses").Use(middleware.OperationRecord())
	expensesRouterWithoutRecord := Router.Group("expenses")
	var expensesApi = v1.ApiGroupApp.FinancialApiGroup.ExpensesApi
	{
		expensesRouter.POST("createExpenses", expensesApi.CreateExpenses)           // 新建Expenses
		expensesRouter.POST("deleteExpenses", expensesApi.DeleteExpenses)           // 删除Expenses
		expensesRouter.POST("deleteExpensesByIds", expensesApi.DeleteExpensesByIds) // 批量删除Expenses
		expensesRouter.PUT("updateExpenses", expensesApi.UpdateExpenses)            // 更新Expenses
	}
	{
		expensesRouterWithoutRecord.GET("findExpenses", expensesApi.FindExpenses)              // 根据ID获取Expenses
		expensesRouterWithoutRecord.GET("getExpensesList", expensesApi.GetExpensesList)        // 获取Expenses列表
		expensesRouterWithoutRecord.POST("getExpensesExcel", expensesApi.GetExpensesListExcel) // 获取Expenses列表
	}
}
