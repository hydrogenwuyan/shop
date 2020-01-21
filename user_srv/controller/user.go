package controller

import (
	"context"
	log "github.com/sirupsen/logrus"
	usersrvmodel "project/shop/user_srv/model"
	usersrvproto "project/shop/user_srv/proto"
)

type Service struct {
}

func (s *Service) QueryUserInfoByUserId(ctx context.Context, request *usersrvproto.CSUserInfo, response *usersrvproto.SCUserInfo) (err error) {
	user, err := usersrvmodel.GetUserService().QueryUserInfoByUserId(request.UserId)
	if err != nil {
		log.WithFields(log.Fields{
			"userId": request.UserId,
			"error":  err,
		}).Error("usersrv: 数据库错误")
		response.Error = &usersrvproto.Error{
			Code:   400,
			Detail: err.Error(),
		}
		return
	}
	if user.Id == 0 {
		log.WithFields(log.Fields{
			"userId": request.UserId,
		}).Warn("usersrv: 用户不存在")
		response.Error = &usersrvproto.Error{
			Code:   404,
			Detail: "user not exist",
		}
		return
	}

	response.Info = &usersrvproto.UserInfo{
		UserId:     user.Id,
		CreateTime: user.CreateTime,
	}
	response.Error = &usersrvproto.Error{
		Code: 200,
	}

	return
}
