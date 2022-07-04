package inventoryManage

import (
	"github.com/tongxz/xs-admin-vue/server/global"
	"github.com/tongxz/xs-admin-vue/server/model/common/request"
	"github.com/tongxz/xs-admin-vue/server/model/inventoryManage"
	inventoryManageReq "github.com/tongxz/xs-admin-vue/server/model/inventoryManage/request"
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

// GetWarehousing 根据id获取Warehousing记录
// Author [piexlmax](https://github.com/piexlmax)
func (warehousingService *WarehousingService) GetWarehousing(id uint) (warehousing inventoryManage.Warehousing, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&warehousing).Error
	return
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
	err = db.Limit(limit).Offset(offset).Find(&warehousings).Error
	return warehousings, total, err
}
