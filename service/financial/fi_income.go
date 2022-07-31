package financial

import (
	"github.com/tongxz/xs-admin-vue/global"
	"github.com/tongxz/xs-admin-vue/model/common/request"
	"github.com/tongxz/xs-admin-vue/model/financial"
	financialReq "github.com/tongxz/xs-admin-vue/model/financial/request"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IncomeService struct {
}

// CreateIncome 创建Income记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeService *IncomeService) CreateIncome(income financial.Income) (err error) {
	err = global.GVA_DB.Create(&income).Error
	return err
}

// DeleteIncome 删除Income记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeService *IncomeService) DeleteIncome(income financial.Income) (err error) {
	err = global.GVA_DB.Delete(&income).Error
	return err
}

// DeleteIncomeByIds 批量删除Income记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeService *IncomeService) DeleteIncomeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]financial.Income{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateIncome 更新Income记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeService *IncomeService) UpdateIncome(income financial.Income) (err error) {
	var oldIncome financial.Income
	upDateMap := make(map[string]interface{})
	upDateMap["incomeData"] = income.IncomeData
	upDateMap["name"] = income.Name
	upDateMap["mobile"] = income.Mobile
	upDateMap["amount"] = income.Amount
	upDateMap["department"] = income.Department
	upDateMap["category"] = income.Category
	upDateMap["payment"] = income.Payment
	upDateMap["invoice"] = income.Invoice
	upDateMap["bill"] = income.Bill
	upDateMap["waiter"] = income.Waiter
	upDateMap["note"] = income.Note
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		db := tx.Where("id = ?", income.ID).Find(&oldIncome)
		txErr := tx.Unscoped().Delete(&financial.IncomeDetail{}, "income_id = ?", income.ID).Error
		if txErr != nil {
			global.GVA_LOG.Debug(txErr.Error())
			return txErr
		}
		global.GVA_LOG.Info("IncomeDetail", zap.Any("income_detail", income.IncomeDetails))
		if len(income.IncomeDetails) > 0 {
			for k := range income.IncomeDetails {
				income.IncomeDetails[k].IncomeID = income.ID
			}
			txErr = tx.Create(&income.IncomeDetails).Error
			if txErr != nil {
				global.GVA_LOG.Debug(txErr.Error())
				return txErr
			}
		}
		txErr = db.Updates(upDateMap).Error
		if txErr != nil {
			global.GVA_LOG.Debug(txErr.Error())
			return txErr
		}
		return nil
	})
	//global.GVA_LOG.Info("incomes:", zap.Any("incomes", income))
	return err
}

// GetIncome 根据id获取Income记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeService *IncomeService) GetIncome(id uint) (income financial.Income, err error) {
	err = global.GVA_DB.Preload("IncomeDetails").Where("id = ?", id).First(&income).Error
	return
}

// GetIncomeInfoList 分页获取Income记录
// Author [piexlmax](https://github.com/piexlmax)
func (incomeService *IncomeService) GetIncomeInfoList(info financialReq.IncomeSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&financial.Income{})
	var incomes []financial.Income
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Mobile != nil {
		db = db.Where("mobile = ?", info.Mobile)
	}
	if info.Invoice != nil {
		db = db.Where("invoice = ?", info.Invoice)
	}
	if info.Payment != nil {
		db = db.Where("payment = ?", info.Payment)
	}
	if info.Category != nil {
		db = db.Where("category = ?", info.Category)
	}
	if info.Waiter != "" {
		db = db.Where("waiter LIKE ?", "%"+info.Waiter+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("incomeData desc").Limit(limit).Offset(offset).Find(&incomes).Error
	return incomes, total, err
}
