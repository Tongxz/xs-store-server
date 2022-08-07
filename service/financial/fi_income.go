package financial

import (
	"fmt"
	"github.com/tongxz/xs-admin-vue/global"
	"github.com/tongxz/xs-admin-vue/model/common/request"
	"github.com/tongxz/xs-admin-vue/model/financial"
	financialReq "github.com/tongxz/xs-admin-vue/model/financial/request"
	"github.com/tongxz/xs-admin-vue/model/system"
	"github.com/xuri/excelize/v2"
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

func (incomeService *IncomeService) GetIncomeList(info financial.Income) (list []financial.Income, err error) {
	// 创建db
	db := global.GVA_DB.Model(&financial.Income{})
	var expensess []financial.Income
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
	err = db.Order("incomeData desc").Find(&expensess).Error
	return expensess, err
}

func (exa *IncomeService) ParseInfoList2Excel(infoList []financial.Income, filePath string) error {
	excel := excelize.NewFile()
	//IncomeData    string         `json:"incomeData" form:"incomeData" gorm:"column:incomeData;comment:收入日期;"`
	//Name          string         `json:"name" form:"name" gorm:"column:name;comment:姓名;"`
	//Mobile        *int           `json:"mobile" form:"mobile" gorm:"column:mobile;comment:手机号码;size:11;"`
	//Amount        *float64       `json:"amount" form:"amount" gorm:"column:amount;comment:金额;"`
	//Department    string         `json:"department" form:"department" gorm:"column:department;comment:部门;"`
	//Category      *int           `json:"category" form:"category" gorm:"column:category;comment:类别;"`
	//Payment       *int           `json:"payment" form:"payment" gorm:"column:payment;comment:收款方式;"`
	//Invoice       *bool          `json:"invoice" form:"invoice" gorm:"column:invoice;comment:是否开票;"`
	//Bill          string         `json:"bill" form:"bill" gorm:"column:bill;comment:发票号;"`
	//Waiter        string         `json:"waiter" form:"waiter" gorm:"column:waiter;comment:负责人;"`
	//Note          string         `json:"note" form:"note" gorm:"column:note;comment:备注;"`
	//IncomeDetails []IncomeDetail `json:"incomeDetails"`
	var arry []financial.Income

	excel.SetSheetRow("Sheet1", "A1", &[]string{"ID", "收入日期", "姓名", "手机号码", "金额", "部门", "类别", "收款方式", "是否开票", "发票号", "负责人", "备注"})
	for i, v := range infoList {
		arry = append(arry, v)
		sysDictionary := system.SysDictionary{}
		myPayment := system.SysDictionary{}
		global.GVA_DB.Where("type = ? OR id = ? and status = ?", v.Department, nil, true).Preload("SysDictionaryDetails", "status = ?", true).First(&sysDictionary)
		global.GVA_DB.Where("type = ? OR id = ? and status = ?", "pay_by", nil, true).Preload("SysDictionaryDetails", "status = ?", true).First(&myPayment)
		for _, dict := range sysDictionary.SysDictionaryDetails {
			if dict.Value == *v.Category {
				arry[i].CategoryName = dict.Label
			}
		}
		for _, dict := range myPayment.SysDictionaryDetails {
			if dict.Value == *v.Payment {
				arry[i].PaymentName = dict.Label
			}
		}
		if v.Department == "food" {
			arry[i].Department = "餐饮部"
		} else if v.Department == "tea" {
			arry[i].Department = "茶饮部"
		} else {
			arry[i].Department = "其他"
		}
		if *v.Invoice {
			arry[i].InvoiceDesc = "是"
		} else {
			arry[i].InvoiceDesc = "否"
		}
	}
	for i, Ware := range arry {
		axis := fmt.Sprintf("A%d", i+2)
		excel.SetSheetRow("Sheet1", axis, &[]interface{}{
			Ware.ID,
			Ware.IncomeData,
			Ware.Name,
			*Ware.Mobile,
			*Ware.Amount,
			Ware.Department,
			Ware.CategoryName,
			Ware.PaymentName,
			Ware.InvoiceDesc,
			Ware.Bill,
			Ware.Waiter,
			Ware.Note,
		})
	}
	err := excel.SaveAs(filePath)
	return err
}
