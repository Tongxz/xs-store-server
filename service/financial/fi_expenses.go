package financial

import (
	"fmt"
	"github.com/tongxz/xs-admin-vue/global"
	"github.com/tongxz/xs-admin-vue/model/common/request"
	"github.com/tongxz/xs-admin-vue/model/financial"
	financialReq "github.com/tongxz/xs-admin-vue/model/financial/request"
	"github.com/xuri/excelize/v2"
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
	err = db.Order("expnDate desc").Limit(limit).Offset(offset).Find(&expensess).Error
	return expensess, total, err
}

// GetExpensesInfoList 分页获取Expenses记录
// Author [piexlmax](https://github.com/piexlmax)
func (expensesService *ExpensesService) GetExpensesList(info financial.Expenses) (list []financial.Expenses, err error) {
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
	err = db.Order("expnDate desc").Find(&expensess).Error
	return expensess, err
}

func (exa *ExpensesService) ParseInfoList2Excel(infoList []financial.Expenses, filePath string) error {
	excel := excelize.NewFile()
	//global.GVA_MODEL
	//ExpnDate   string   `json:"expnDate" form:"expnDate" gorm:"column:expnDate;comment:收入日期;"`
	//Content    string   `json:"content" form:"content" gorm:"column:content;comment:支出内容说明;"`
	//Amount     *float64 `json:"amount" form:"amount" gorm:"column:amount;comment:支出金额;"`
	//Department string   `json:"department" form:"department" gorm:"column:department;comment:部门;"`
	//Type       *int     `json:"type" form:"type" gorm:"column:type;comment:支出类型;"`
	//Payment    *int     `json:"payment" form:"payment" gorm:"column:payment;comment:支付方式;"`
	//Executor   string   `json:"executor" form:"executor" gorm:"column:executor;comment:执行该操作的人员;"`
	//Invoice    *bool    `json:"invoice" form:"invoice" gorm:"column:invoice;comment:是否开票;"`
	//Note       string   `json:"note" form:"note" gorm:"column:note;comment:备注说明;"`
	excel.SetSheetRow("Sheet1", "A1", &[]string{"ID", "收入日期", "支出说明", "支出金额", "部门", "支出类型", "支付方式", "操作人员", "是否开票", "备注说明"})
	for i, Ware := range infoList {
		axis := fmt.Sprintf("A%d", i+2)
		excel.SetSheetRow("Sheet1", axis, &[]interface{}{
			Ware.ID,
			Ware.ExpnDate,
			Ware.Content,
			*Ware.Amount,
			Ware.Department,
			Ware.Type,
			Ware.Payment,
			Ware.Executor,
			Ware.Invoice,
			Ware.Note,
		})
	}
	err := excel.SaveAs(filePath)
	return err
}
