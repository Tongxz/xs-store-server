package financial

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tongxz/xs-admin-vue/server/global"
	"github.com/tongxz/xs-admin-vue/server/model/common/request"
	"github.com/tongxz/xs-admin-vue/server/model/common/response"
	"github.com/tongxz/xs-admin-vue/server/model/financial"
	financialReq "github.com/tongxz/xs-admin-vue/server/model/financial/request"
	"github.com/tongxz/xs-admin-vue/server/service"
	"github.com/tongxz/xs-admin-vue/server/utils"
	"go.uber.org/zap"
)

type IncomeApi struct {
}

var incomeService = service.ServiceGroupApp.FinancialServiceGroup.IncomeService

// CreateIncome 创建Income
// @Tags Income
// @Summary 创建Income
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body financial.Income true "创建Income"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /income/createIncome [post]
func (incomeApi *IncomeApi) CreateIncome(c *gin.Context) {
	var income financial.Income
	uuid.New()
	_ = c.ShouldBindJSON(&income)
	//if err := utils.Verify(income, utils.IncomeVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	if err := incomeService.CreateIncome(income); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIncome 删除Income
// @Tags Income
// @Summary 删除Income
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body financial.Income true "删除Income"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /income/deleteIncome [delete]
func (incomeApi *IncomeApi) DeleteIncome(c *gin.Context) {
	var income financial.Income
	_ = c.ShouldBindJSON(&income)
	if err := incomeService.DeleteIncome(income); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIncomeByIds 批量删除Income
// @Tags Income
// @Summary 批量删除Income
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Income"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /income/deleteIncomeByIds [delete]
func (incomeApi *IncomeApi) DeleteIncomeByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := incomeService.DeleteIncomeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIncome 更新Income
// @Tags Income
// @Summary 更新Income
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body financial.Income true "更新Income"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /income/updateIncome [put]
func (incomeApi *IncomeApi) UpdateIncome(c *gin.Context) {
	var income financial.Income
	_ = c.ShouldBindJSON(&income)
	if err := utils.Verify(income, utils.IncomeVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := incomeService.UpdateIncome(income); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIncome 用id查询Income
// @Tags Income
// @Summary 用id查询Income
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query financial.Income true "用id查询Income"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /income/findIncome [get]
func (incomeApi *IncomeApi) FindIncome(c *gin.Context) {
	var income financial.Income
	_ = c.ShouldBindQuery(&income)
	if reincome, err := incomeService.GetIncome(income.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reincome": reincome}, c)
	}
}

// GetIncomeList 分页获取Income列表
// @Tags Income
// @Summary 分页获取Income列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query financialReq.IncomeSearch true "分页获取Income列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /income/getIncomeList [get]
func (incomeApi *IncomeApi) GetIncomeList(c *gin.Context) {
	var pageInfo financialReq.IncomeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := incomeService.GetIncomeInfoList(pageInfo); err != nil {
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
