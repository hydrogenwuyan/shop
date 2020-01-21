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
	s    *userService
)

// 用户接口
type UserService interface {
	QueryUserInfoByUserId(id int64) (user *UserEntity, err error)
}

type userService struct {
	db *gorm.DB
}

func (s *userService) QueryUserInfoByUserId(userId int64) (user *UserEntity, err error) {
	user = &UserEntity{}
	sql := fmt.Sprintf("select %s,%s,%s,%s,%s,%s from %s where %s=?",
		User_Id, User_UserId, User_Password, User_CreateTime, User_UpdateTime, User_DeleteTime, User_TableName, User_UserId)
	err = s.db.Raw(sql, userId).Scan(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = nil
			return
		}

		log.Fatalf("usersrv: 数据库出错, error: %v", err)
		return
	}

	return
}

func Init() {
	once.Do(func() {
		s = &userService{
			db: commondb.GetDB().DB(),
		}
	})
}

func GetUserService() UserService {
	return s
}
