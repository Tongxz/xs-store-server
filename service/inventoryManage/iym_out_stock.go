package inventoryManage

import (
	"github.com/tongxz/xs-admin-vue/server/global"
	"github.com/tongxz/xs-admin-vue/server/model/common/request"
	"github.com/tongxz/xs-admin-vue/server/model/inventoryManage"
	inventoryManageReq "github.com/tongxz/xs-admin-vue/server/model/inventoryManage/request"
)

type OutStockService struct {
}

// CreateOutStock 创建OutStock记录
// Author [piexlmax](https://github.com/piexlmax)
func (stockService *OutStockService) CreateOutStock(stock inventoryManage.OutStock) (err error) {
	err = global.GVA_DB.Create(&stock).Error
	return err
}

// DeleteOutStock 删除OutStock记录
// Author [piexlmax](https://github.com/piexlmax)
func (stockService *OutStockService) DeleteOutStock(stock inventoryManage.OutStock) (err error) {
	err = global.GVA_DB.Delete(&stock).Error
	return err
}

// DeleteOutStockByIds 批量删除OutStock记录
// Author [piexlmax](https://github.com/piexlmax)
func (stockService *OutStockService) DeleteOutStockByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]inventoryManage.OutStock{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOutStock 更新OutStock记录
// Author [piexlmax](https://github.com/piexlmax)
func (stockService *OutStockService) UpdateOutStock(stock inventoryManage.OutStock) (err error) {
	err = global.GVA_DB.Save(&stock).Error
	return err
}

// GetOutStock 根据id获取OutStock记录
// Author [piexlmax](https://github.com/piexlmax)
func (stockService *OutStockService) GetOutStock(id uint) (stock inventoryManage.OutStock, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&stock).Error
	return
}

// GetOutStockInfoList 分页获取OutStock记录
// Author [piexlmax](https://github.com/piexlmax)
func (stockService *OutStockService) GetOutStockInfoList(info inventoryManageReq.OutStockSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&inventoryManage.OutStock{})
	var stocks []inventoryManage.OutStock
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Type != nil {
		db = db.Where("type = ?", info.Type)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.Item_type != "" {
		db = db.Where("item_type = ?", info.Item_type)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&stocks).Error
	return stocks, total, err
}
