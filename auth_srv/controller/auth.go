package controller

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"project/shop/auth_srv/model/token"
	authsrvproto "project/shop/auth_srv/proto"
	"strconv"
)

type Service struct {
}

// 设置token
func (s *Service) SetTokenByUserId(ctx context.Context, req *authsrvproto.CSTokenSet, rsp *authsrvproto.SCTokenSet) error {
	token, err := token.GetTokenService().SetToken(&token.Subject{
		Id: strconv.FormatInt(req.UserId, 10),
	})

	if err != nil {
		rsp.Error = &authsrvproto.Error{
			Detail: err.Error(),
		}

		log.WithFields(log.Fields{
			"error": err,
		}).Error("authsrv:  token生成失败")

		return err
	}

	rsp.Error = &authsrvproto.Error{
		Code: 200,
	}
	rsp.UserId = req.UserId
	rsp.Token = token

	return nil
}

// 清除用户token
func (s *Service) ClearTokenByUserId(ctx context.Context, req *authsrvproto.CSTokenClear, rsp *authsrvproto.SCTokenClear) error {
	log.WithFields(log.Fields{
		"CSTokenClear": *req,
	}).Debug("authsrv: 收到清除用户token请求")

	err := token.GetTokenService().ClearToken(req.Token)
	if err != nil {
		rsp.Error = &authsrvproto.Error{
			Detail: err.Error(),
		}

		log.WithFields(log.Fields{
			"error": err,
		}).Error("authsrv:  清除用户token失败")

		return err
	}

	rsp.Error = &authsrvproto.Error{
		Code: 200,
	}

	return nil
}

// 请求uid
func (s *Service) GetUserIdByToken(ctx context.Context, req *authsrvproto.CSUserIdGet, rsp *authsrvproto.SCUserIdGet) error {
	log.WithFields(log.Fields{
		"CSTokenGet": *req,
	}).Debug("authsrv: 请求uid")

	userId, err := token.GetTokenService().GetUserId(req.Token)
	if err != nil {
		rsp.Error = &authsrvproto.Error{
			Detail: err.Error(),
		}

		log.WithFields(log.Fields{
			"error": err,
		}).Error("authsrv:  获取缓存的userId失败")

		return err
	}

	fmt.Println(userId)

	rsp.Error = &authsrvproto.Error{
		Code: 200,
	}
	rsp.UserId = userId

	return nil
}
