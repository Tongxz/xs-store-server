package financial

import (
	"github.com/gin-gonic/gin"
	"github.com/tongxz/xs-admin-vue/global"
	"github.com/tongxz/xs-admin-vue/model/common/request"
	"github.com/tongxz/xs-admin-vue/model/common/response"
	"github.com/tongxz/xs-admin-vue/model/financial"
	financialReq "github.com/tongxz/xs-admin-vue/model/financial/request"
	"github.com/tongxz/xs-admin-vue/service"
	"go.uber.org/zap"
	"strings"
)

type ExpensesApi struct {
}

var expensesService = service.ServiceGroupApp.FinancialServiceGroup.ExpensesService

// CreateExpenses 创建Expenses
// @Tags Expenses
// @Summary 创建Expenses
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body financial.Expenses true "创建Expenses"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /expenses/createExpenses [post]
func (expensesApi *ExpensesApi) CreateExpenses(c *gin.Context) {
	var expenses financial.Expenses
	_ = c.ShouldBindJSON(&expenses)
	if err := expensesService.CreateExpenses(expenses); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteExpenses 删除Expenses
// @Tags Expenses
// @Summary 删除Expenses
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body financial.Expenses true "删除Expenses"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /expenses/deleteExpenses [delete]
func (expensesApi *ExpensesApi) DeleteExpenses(c *gin.Context) {
	var expenses financial.Expenses
	_ = c.ShouldBindJSON(&expenses)
	if err := expensesService.DeleteExpenses(expenses); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteExpensesByIds 批量删除Expenses
// @Tags Expenses
// @Summary 批量删除Expenses
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Expenses"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /expenses/deleteExpensesByIds [delete]
func (expensesApi *ExpensesApi) DeleteExpensesByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := expensesService.DeleteExpensesByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateExpenses 更新Expenses
// @Tags Expenses
// @Summary 更新Expenses
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body financial.Expenses true "更新Expenses"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /expenses/updateExpenses [put]
func (expensesApi *ExpensesApi) UpdateExpenses(c *gin.Context) {
	var expenses financial.Expenses
	_ = c.ShouldBindJSON(&expenses)
	if err := expensesService.UpdateExpenses(expenses); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindExpenses 用id查询Expenses
// @Tags Expenses
// @Summary 用id查询Expenses
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query financial.Expenses true "用id查询Expenses"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /expenses/findExpenses [get]
func (expensesApi *ExpensesApi) FindExpenses(c *gin.Context) {
	var expenses financial.Expenses
	_ = c.ShouldBindQuery(&expenses)
	if reexpenses, err := expensesService.GetExpenses(expenses.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reexpenses": reexpenses}, c)
	}
}

// GetExpensesList 分页获取Expenses列表
// @Tags Expenses
// @Summary 分页获取Expenses列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query financialReq.ExpensesSearch true "分页获取Expenses列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /expenses/getExpensesList [get]
func (expensesApi *ExpensesApi) GetExpensesList(c *gin.Context) {
	var pageInfo financialReq.ExpensesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := expensesService.GetExpensesInfoList(pageInfo); err != nil {
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

// GetExpensesListExcel 获取Expenses列表excel
// @Tags Expenses
// @Summary 分页获取Expenses列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query financialReq.ExpensesSearch true "分页获取Expenses列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /expenses/getExpensesList [get]
func (expensesApi *ExpensesApi) GetExpensesListExcel(c *gin.Context) {
	var pageInfo financial.ExpensesExcel
	_ = c.ShouldBindJSON(&pageInfo)
	if strings.Index(pageInfo.FileName, "..") > -1 {
		response.FailWithMessage("包含非法字符", c)
		return
	}
	filePath := global.GVA_CONFIG.Excel.Dir + pageInfo.FileName
	if list, err := expensesService.GetExpensesList(pageInfo.InfoList); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		err := expensesService.ParseInfoList2Excel(list, filePath)
		if err != nil {
			global.GVA_LOG.Error("转换Excel失败!", zap.Error(err))
			response.FailWithMessage("转换Excel失败", c)
			return
		}
		c.Writer.Header().Add("success", "true")
		c.File(filePath)
	}
}
