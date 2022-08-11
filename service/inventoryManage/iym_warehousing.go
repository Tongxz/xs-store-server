package inventoryManage

import (
	"fmt"
	"github.com/tongxz/xs-admin-vue/global"
	"github.com/tongxz/xs-admin-vue/model/common/request"
	"github.com/tongxz/xs-admin-vue/model/inventoryManage"
	inventoryManageReq "github.com/tongxz/xs-admin-vue/model/inventoryManage/request"
	"github.com/xuri/excelize/v2"
)

type WarehousingService struct {
}

// CreateWarehousing 创建Warehousing记录
// Author [piexlmax](https://github.com/piexlmax)
func (warehousingService *WarehousingService) CreateWarehousing(warehousing inventoryManage.Warehousing) (err error) {
	err = global.GVA_DB.Create(&warehousing).Error
	return err
}

// DeleteWarehousing 删除Warehousing记录
// Author [piexlmax](https://github.com/piexlmax)
func (warehousingService *WarehousingService) DeleteWarehousing(warehousing inventoryManage.Warehousing) (err error) {
	err = global.GVA_DB.Delete(&warehousing).Error
	return err
}

// DeleteWarehousingByIds 批量删除Warehousing记录
// Author [piexlmax](https://github.com/piexlmax)
func (warehousingService *WarehousingService) DeleteWarehousingByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]inventoryManage.Warehousing{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateWarehousing 更新Warehousing记录
// Author [piexlmax](https://github.com/piexlmax)
func (warehousingService *WarehousingService) UpdateWarehousing(warehousing inventoryManage.Warehousing) (err error) {
	err = global.GVA_DB.Save(&warehousing).Error
	return err
}

// UpdateWarehousingMargin 根据id更新WarehousingMargin记录
// Author [piexlmax](https://github.com/piexlmax)
func (warehousingService *WarehousingService) UpdateWarehousingMargin(warehousing inventoryManage.Warehousing) (err error) {
	var old inventoryManage.Warehousing
	err = global.GVA_DB.Where("id = ?", warehousing.ID).First(&old).Error
	err = global.GVA_DB.Model(&warehousing).Where("id = ?", warehousing.ID).Updates(map[string]interface{}{"margin": *old.Margin + *warehousing.Margin, "quantity": *old.Quantity + *warehousing.Margin}).Error
	return err
}

// GetWarehousing 根据id获取Warehousing记录
// Author [piexlmax](https://github.com/piexlmax)
func (warehousingService *WarehousingService) GetWarehousing(id uint) (warehousing inventoryManage.Warehousing, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&warehousing).Error
	return
}
func (warehousingService *WarehousingService) GetWarehousingName(incomeType *int) (list interface{}, err error) {
	// 创建db
	db := global.GVA_DB.Model(&inventoryManage.WarehousingName{})
	var warehousings []inventoryManage.WarehousingName
	// 如果有条件搜索 下方会自动创建搜索语句
	if incomeType != nil {
		db = db.Where("income_type = ?", incomeType)
	}
	err = db.Where("Margin > ? ", 0).Find(&warehousings).Error
	return warehousings, err
}

// GetWarehousingInfoList 分页获取Warehousing记录
// Author [piexlmax](https://github.com/piexlmax)
func (warehousingService *WarehousingService) GetWarehousingInfoList(info inventoryManageReq.WarehousingSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&inventoryManage.Warehousing{})
	var warehousings []inventoryManage.Warehousing
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Department != "" {
		db = db.Where("department = ?", info.Department)
	}
	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	if info.Payment != nil {
		db = db.Where("payment = ?", info.Payment)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("created_at desc").Limit(limit).Offset(offset).Find(&warehousings).Error
	return warehousings, total, err
}

func (warehousingService *WarehousingService) ParseInfoList2Excel(infoList []inventoryManage.Warehousing, filePath string) error {
	excel := excelize.NewFile()
	excel.SetSheetRow("Sheet1", "A1", &[]string{"ID", "物品图片", "物品名称", "所属部门", "所属分类", "收入分类", "支付方式", "物品数量", "剩余物品数量", "物品单位", "物品单价", "成本价", "总金额", "入库备注/说明"})
	for i, Ware := range infoList {
		axis := fmt.Sprintf("A%d", i+2)
		excel.SetSheetRow("Sheet1", axis, &[]interface{}{
			Ware.ID,
			Ware.ImgUrl,
			Ware.Name,
			Ware.Department,
			Ware.Type,
			Ware.IncomeType,
			Ware.Payment,
			*Ware.Quantity,
			*Ware.Margin,
			Ware.Unit,
			*Ware.UnitPrice,
			*Ware.Cost,
			*Ware.Amount,
			Ware.Remarks,
		})
	}
	err := excel.SaveAs(filePath)
	return err
}
