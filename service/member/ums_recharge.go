package member

import (
	"github.com/tongxz/xs-admin-vue/global"
	"github.com/tongxz/xs-admin-vue/model/common/request"
	"github.com/tongxz/xs-admin-vue/model/member"
	memberReq "github.com/tongxz/xs-admin-vue/model/member/request"
	"gorm.io/gorm"
)

type RechargeService struct {
}

// CreateRecharge 创建Recharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (rechargeService *RechargeService) CreateRecharge(recharge member.Recharge) (err error) {
	var old member.Member
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Create(&recharge).Error
		err = tx.Where("id = ?", recharge.MemberId).First(&old).Error
		err = tx.Model(&old).Where("id = ?", recharge.MemberId).Updates(map[string]interface{}{"member_balance": *old.MemberBalance + *recharge.Amount}).Error
		return err
	})
	return err
}

// DeleteRecharge 删除Recharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (rechargeService *RechargeService) DeleteRecharge(recharge member.Recharge) (err error) {
	err = global.GVA_DB.Delete(&recharge).Error
	return err
}

// DeleteRechargeByIds 批量删除Recharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (rechargeService *RechargeService) DeleteRechargeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]member.Recharge{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateRecharge 更新Recharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (rechargeService *RechargeService) UpdateRecharge(recharge member.Recharge) (err error) {
	err = global.GVA_DB.Save(&recharge).Error
	return err
}

// GetRecharge 根据id获取Recharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (rechargeService *RechargeService) GetRecharge(id uint) (recharge member.Recharge, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&recharge).Error
	return
}

// GetRechargeInfoList 分页获取Recharge记录
// Author [piexlmax](https://github.com/piexlmax)
func (rechargeService *RechargeService) GetRechargeInfoList(info memberReq.RechargeSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&member.Recharge{})
	var recharges []member.Recharge
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.MemberId != nil {
		db = db.Where("member_id = ?", info.MemberId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&recharges).Error
	return recharges, total, err
}
