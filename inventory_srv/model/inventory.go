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

const (
	MaxRetryTimes = 10 // 最大重试次数
)

// 库存接口
type InventoryService interface {
	// 通过商品id查询库存
	QueryInventoryInfoByShopId(id int64) (entity *InventoryEntity, err error)
	// 更新库存，库存减少
	UpdateAboutBuy(shopId int64, num int64) (money int64, err error)
}

type inventoryService struct {
	db *gorm.DB
}

func (s *inventoryService) QueryInventoryInfoByShopId(shopId int64) (entity *InventoryEntity, err error) {
	return s.queryInventoryInfoByShopId(shopId)
}

func (s *inventoryService) queryInventoryInfoByShopId(shopId int64) (entity *InventoryEntity, err error) {
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

func (s *inventoryService) UpdateAboutBuy(shopId int64, num int64) (money int64, err error) {
	// 开启事务
	tx := s.db.Begin()
	defer func() {
		// 出错回滚
		if r := recover(); r != nil {
			tx.Rollback()
			log.Errorf("inventorysrv: 数据库出错")
		}
	}()

	if tx.Error != nil {
		err = tx.Error
		return
	}

	var times int

	var updateInventoryFunc func() error
	updateInventoryFunc = func() (err2 error) {
		// 查询库存
		inventory, err2 := s.queryInventoryInfoByShopId(shopId)
		if err2 != nil {
			return
		}

		// 判断库存是否足够
		if inventory.Num < num {
			err2 = fmt.Errorf("库存不足")
			return
		}

		// 减少库存
		sql := fmt.Sprintf("update %s set %s=%s-?,%s=%s+1 where %s=? and %s=?", Inventory_TableName, Inventory_Num, Inventory_Num, Inventory_Version, Inventory_Version, Inventory_ShopId, Inventory_Version)
		err2 = tx.Exec(sql, num, shopId, inventory.Version).Error
		if err2 != nil {
			if err2 == gorm.ErrRecordNotFound {
				times++
				if times > MaxRetryTimes {
					return
				}
				// 重试
				updateInventoryFunc()
			}

			return
		}

		money = inventory.Money
		return
	}

	// 刷新数据库
	err = updateInventoryFunc()
	if err != nil {
		return
	}

	// 提交事务
	err = tx.Commit().Error

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
