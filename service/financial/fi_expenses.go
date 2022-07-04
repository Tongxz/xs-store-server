package financial

import (
	"github.com/tongxz/xs-admin-vue/server/global"
	"github.com/tongxz/xs-admin-vue/server/model/common/request"
	"github.com/tongxz/xs-admin-vue/server/model/financial"
	financialReq "github.com/tongxz/xs-admin-vue/server/model/financial/request"
)

type ExpensesService struct {
}

// CreateExpenses 创建Expenses记录
// Author [piexlmax](https://github.com/piexlmax)
func (expensesService *ExpensesService) CreateExpenses(expenses financial.Expenses) (err error) {
	err = global.GVA_DB.Create(&expenses).Error
	return err
}

// DeleteExpenses 删除Expenses记录
// Author [piexlmax](https://github.com/piexlmax)
func (expensesService *ExpensesService) DeleteExpenses(expenses financial.Expenses) (err error) {
	err = global.GVA_DB.Delete(&expenses).Error
	return err
}

// DeleteExpensesByIds 批量删除Expenses记录
// Author [piexlmax](https://github.com/piexlmax)
func (expensesService *ExpensesService) DeleteExpensesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]financial.Expenses{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateExpenses 更新Expenses记录
// Author [piexlmax](https://github.com/piexlmax)
func (expensesService *ExpensesService) UpdateExpenses(expenses financial.Expenses) (err error) {
	err = global.GVA_DB.Save(&expenses).Error
	return err
}

// GetExpenses 根据id获取Expenses记录
// Author [piexlmax](https://github.com/piexlmax)
func (expensesService *ExpensesService) GetExpenses(id uint) (expenses financial.Expenses, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&expenses).Error
	return
}

// GetExpensesInfoList 分页获取Expenses记录
// Author [piexlmax](https://github.com/piexlmax)
func (expensesService *ExpensesService) GetExpensesInfoList(info financialReq.ExpensesSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&financial.Expenses{})
	var expensess []financial.Expenses
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	if info.Payment != nil {
		db = db.Where("payment = ?", info.Payment)
	}
	if info.Executor != "" {
		db = db.Where("executor = ?", info.Executor)
	}
	if info.Invoice != nil {
		db = db.Where("invoice = ?", info.Invoice)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&expensess).Error
	return expensess, total, err
}
