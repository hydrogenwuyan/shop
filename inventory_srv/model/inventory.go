package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	commondb "project/shop/common/db"
	"sync"
)

var (
	once sync.Once
	s    *inventoryService
)

// 库存接口
type InventoryService interface {
	// 通过商品id查询库存
	QueryInventoryInfoByShopId(id int64) (entity *InventoryEntity, err error)
	// 更新库存，库存减一
	UpdateByShopId(shopId int64, version int64) (err error)
}

type inventoryService struct {
	db *gorm.DB
}

func (s *inventoryService) QueryInventoryInfoByShopId(shopId int64) (entity *InventoryEntity, err error) {
	entity = &InventoryEntity{}
	err = s.db.Find(entity, "shopId=?", shopId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			return
		}

		log.Errorf("inventorysrv: 数据库出错, error: %v", err)
		return
	}

	return
}

func (s *inventoryService) UpdateByShopId(shopId int64, version int64) (err error) {
	sql := fmt.Sprintf("update %s set %s=%s-1 where %s=? and %s=?", Inventory_TableName, Inventory_Num, Inventory_Num, Inventory_ShopId, Inventory_Version)
	err = s.db.Raw(sql, shopId, version).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			return
		}

		log.Errorf("inventorysrv: 数据库出错, error: %v", err)
		return
	}

	return
}

func Init() {
	once.Do(func() {
		s = &inventoryService{
			db: commondb.GetDB().DB(),
		}
	})
}

func GetInventoryService() InventoryService {
	return s
}
