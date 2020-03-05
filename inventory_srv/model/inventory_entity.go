package model

// 库存数据
type InventoryEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	ShopId     int64 `gorm:"column:shopId"`
	Num        int64 `gorm:"column:num"`
	Money      int64 `gorm:"column:money"`
	Version    int64 `gorm:"column:version"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *InventoryEntity) GetId() int64 {
	return e.Id
}

func (e *InventoryEntity) TableName() string {
	return "t_inventory_entity"
}

const (
	Inventory_TableName  = "t_inventory_entity"
	Inventory_Id         = "id"
	Inventory_ShopId     = "shopId"
	Inventory_Num        = "num"
	Inventory_Version    = "version"
	Inventory_UpdateTime = "updateTime"
	Inventory_CreateTime = "createTime"
	Inventory_DeleteTime = "deleteTime"
)
