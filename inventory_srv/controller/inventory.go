package controller

import (
	"context"
	inventorysrvproto "project/shop/inventory_srv/proto"
)

type Service struct {
}

// 查询库存
func (s *Service) QueryInventoryInfoByShopId(ctx context.Context, request *inventorysrvproto.CSInventoryInfo, response *inventorysrvproto.SCInventoryInfo) (err error) {
	return
}

// 更新库存
func (s *Service) InventoryBuy(ctx context.Context, request *inventorysrvproto.CSInventoryBuy, response *inventorysrvproto.SCInventoryBuy) (err error) {
	return
}
