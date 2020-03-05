package controller

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix"
	log "github.com/sirupsen/logrus"
	inventorysrvproto "project/shop/inventory_srv/proto"
	ordersrvmodel "project/shop/order_srv/model"
	ordersrvproto "project/shop/order_srv/proto"
	"time"
)

var (
	inventorysrvClient inventorysrvproto.InventoryService // 库存服务client
	ordersrvClient     ordersrvproto.OrderService         // 用户服务client
)

// 初始化
func Init() {
	cl := hystrix.NewClientWrapper()(client.DefaultClient)
	ordersrvClient = ordersrvproto.NewOrderService("shop.order.srv", cl)
	inventorysrvClient = inventorysrvproto.NewInventoryService("shop.inventory.srv", cl)
}

type Service struct {
}

// 创建订单
func (s *Service) CreateOrder(ctx context.Context, request *ordersrvproto.CSOrderCreate, response *ordersrvproto.SCOrderCreate) (err error) {
	shopId := request.ShopId
	num := request.Num

	// 锁定库存中的商品
	scInventoryInfo, err := inventorysrvClient.InventoryBuy(ctx, &inventorysrvproto.CSInventoryBuy{
		ShopId: shopId,
		Num:    num,
	})
	if err != nil {
		log.Errorf("ordersrv: rpc请求错误, error: %v", err)
		return
	}

	if scInventoryInfo.Error.Code != 200 {
		log.Errorf("ordersrv: 查询订单请求错误, error: %s", scInventoryInfo.Error.Detail)
		return
	}

	// 创建订单
	order := &ordersrvmodel.OrderEntity{
		Status:     ordersrvmodel.OrderStatusCreate,
		Money:      scInventoryInfo.Money,
		Num:        scInventoryInfo.Num,
		UserId:     request.UserId,
		ShopId:     shopId,
		CreateTime: time.Now().Unix(),
	}
	err = ordersrvmodel.GetOrderService().Update(order)
	if err != nil {
		log.Errorf("ordersrv: 数据库错误, error: %v", err)
		return
	}

	response.Info = &ordersrvproto.OrderInfo{
		OrderId: 1,
		ShopId:  request.ShopId,
		Num:     scInventoryInfo.Num,
		UserId:  request.UserId,
		Status:  ordersrvmodel.OrderStatusCreate,
		Money:   scInventoryInfo.Money,
	}
	response.Error = &ordersrvproto.Error{
		Code: 200,
	}

	return
}

// 确认订单
func (s *Service) ConfirmOrder(ctx context.Context, request *ordersrvproto.CSOrderConfirm, response *ordersrvproto.SCOrderConfirm) (err error) {
	return
}
