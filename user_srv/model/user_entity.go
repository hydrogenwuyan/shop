package model

// 用户数据
type UserEntity struct {
	Id         int64  `gorm:"primary_key;column:id"`
	UserId     int64  `gorm:"column:userId"`
	Password   string `gorm:"column:password"`
	UpdateTime int64  `gorm:"column:updateTime"`
	CreateTime int64  `gorm:"column:createTime"`
	DeleteTime int64  `gorm:"column:deleteTime"`
}

func (e *UserEntity) GetId() int64 {
	return e.Id
}

func (e *UserEntity) TableName() string {
	return "t_user_entity"
}

const (
	User_TableName  = "t_user_entity"
	User_Id         = "id"
	User_UserId     = "userId"
	User_Password   = "password"
	User_UpdateTime = "updateTime"
	User_CreateTime = "createTime"
	User_DeleteTime = "deleteTime"
)
