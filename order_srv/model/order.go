package model

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	commondb "project/shop/common/db"
	"sync"
)

const (
	OrderStatusCreate = iota + 1 // 订单刚创建
	OrderStatusPay               // 订单已付款
	OrderStatusCancel            // 订单已被取消
)

var (
	once sync.Once
	s    *orderService
)

// 订单接口
type OrderService interface {
	// 通过商品id查询订单
	QueryOrderInfoByOrderId(id int64) (order *OrderEntity, err error)
	// 更新订单数据
	Update(order *OrderEntity) (err error)
}

type orderService struct {
	db *gorm.DB
}

func (s *orderService) QueryOrderInfoByOrderId(id int64) (order *OrderEntity, err error) {
	order = &OrderEntity{}
	err = s.db.First(order, "id=?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			return
		}

		log.Fatalf("ordersrv: 数据库出错, error: %v", err)
		return
	}

	return
}

func (s *orderService) Update(order *OrderEntity) (err error) {
	err = s.db.Save(order).Error
	if err != nil {
		log.Fatalf("ordersrv: 数据库出错, error: %v", err)
		return
	}

	return
}

func Init() {
	once.Do(func() {
		s = &orderService{
			db: commondb.GetDB().DB(),
		}
	})
}

func GetOrderService() OrderService {
	return s
}
