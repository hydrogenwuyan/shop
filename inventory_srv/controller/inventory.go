package controller

import (
	"context"
	log "github.com/sirupsen/logrus"
	inventorysrvmodel "project/shop/inventory_srv/model"
	inventorysrvproto "project/shop/inventory_srv/proto"
)

var (
	s *Service
)

type Service struct {
}

// 查询库存
func (s *Service) QueryInventoryInfoByShopId(ctx context.Context, request *inventorysrvproto.CSInventoryInfo, response *inventorysrvproto.SCInventoryInfo) (err error) {
	return
}

// 商品购买，库存减少
func (s *Service) InventoryBuy(ctx context.Context, request *inventorysrvproto.CSInventoryBuy, response *inventorysrvproto.SCInventoryBuy) (err error) {
	money, err := inventorysrvmodel.GetInventoryService().UpdateAboutBuy(request.ShopId, request.Num)
	if err != nil {
		log.Warnf("inventorysrv: 购买商品失败")
		return
	}

	response.Error = &inventorysrvproto.Error{
		Code: 200,
	}
	response.Money = money

	return
}

// 订单取消，库存增加
func (s *Service) InventoryCancel(ctx context.Context, request *inventorysrvproto.CSInventoryCancel, response *inventorysrvproto.SCInventoryCancel) (err error) {
	return
}
