package model

// 订单数据
type OrderEntity struct {
	Id         int64 `gorm:"primary_key;column:id"`
	Status     int32 `gorm:"column:status"`
	ShopId     int64 `gorm:"column:shopId"`
	UserId     int64 `gorm:"column:userId"`
	Num        int64 `gorm:"column:num"`
	Money      int64 `gorm:"column:money"`
	UpdateTime int64 `gorm:"column:updateTime"`
	CreateTime int64 `gorm:"column:createTime"`
	DeleteTime int64 `gorm:"column:deleteTime"`
}

func (e *OrderEntity) GetId() int64 {
	return e.Id
}

func (e *OrderEntity) TableName() string {
	return "t_order_entity"
}

const (
	Order_TableName  = "t_order_entity"
	Order_Id         = "id"
	Order_Status     = "status"
	Order_Money      = "money"
	Order_UpdateTime = "updateTime"
	Order_CreateTime = "createTime"
	Order_DeleteTime = "deleteTime"
)
