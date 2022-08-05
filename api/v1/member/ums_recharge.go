package member

import (
	"github.com/gin-gonic/gin"
	"github.com/tongxz/xs-admin-vue/global"
	"github.com/tongxz/xs-admin-vue/model/common/request"
	"github.com/tongxz/xs-admin-vue/model/common/response"
	"github.com/tongxz/xs-admin-vue/model/member"
	memberReq "github.com/tongxz/xs-admin-vue/model/member/request"
	"github.com/tongxz/xs-admin-vue/service"
	"go.uber.org/zap"
)

type RechargeApi struct {
}

var rechargeService = service.ServiceGroupApp.MemberServiceGroup.RechargeService

// CreateRecharge 创建Recharge
// @Tags Recharge
// @Summary 创建Recharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body member.Recharge true "创建Recharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /recharge/createRecharge [post]
func (rechargeApi *RechargeApi) CreateRecharge(c *gin.Context) {
	var recharge member.Recharge
	_ = c.ShouldBindJSON(&recharge)
	if err := rechargeService.CreateRecharge(recharge); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRecharge 删除Recharge
// @Tags Recharge
// @Summary 删除Recharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body member.Recharge true "删除Recharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /recharge/deleteRecharge [delete]
func (rechargeApi *RechargeApi) DeleteRecharge(c *gin.Context) {
	var recharge member.Recharge
	_ = c.ShouldBindJSON(&recharge)
	if err := rechargeService.DeleteRecharge(recharge); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRechargeByIds 批量删除Recharge
// @Tags Recharge
// @Summary 批量删除Recharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Recharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /recharge/deleteRechargeByIds [delete]
func (rechargeApi *RechargeApi) DeleteRechargeByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := rechargeService.DeleteRechargeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRecharge 更新Recharge
// @Tags Recharge
// @Summary 更新Recharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body member.Recharge true "更新Recharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /recharge/updateRecharge [put]
func (rechargeApi *RechargeApi) UpdateRecharge(c *gin.Context) {
	var recharge member.Recharge
	_ = c.ShouldBindJSON(&recharge)
	if err := rechargeService.UpdateRecharge(recharge); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRecharge 用id查询Recharge
// @Tags Recharge
// @Summary 用id查询Recharge
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query member.Recharge true "用id查询Recharge"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /recharge/findRecharge [get]
func (rechargeApi *RechargeApi) FindRecharge(c *gin.Context) {
	var recharge member.Recharge
	_ = c.ShouldBindQuery(&recharge)
	if rerecharge, err := rechargeService.GetRecharge(recharge.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerecharge": rerecharge}, c)
	}
}

// GetRechargeList 分页获取Recharge列表
// @Tags Recharge
// @Summary 分页获取Recharge列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query memberReq.RechargeSearch true "分页获取Recharge列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /recharge/getRechargeList [get]
func (rechargeApi *RechargeApi) GetRechargeList(c *gin.Context) {
	var pageInfo memberReq.RechargeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := rechargeService.GetRechargeInfoList(pageInfo); err != nil {
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
