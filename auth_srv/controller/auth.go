package controller

import (
	"context"
	"github.com/micro/go-micro/util/log"
	"project/shop/auth_srv/model/token"
	authsrvproto "project/shop/auth_srv/proto"
	"strconv"
)

type Service struct {
}

// 设置token
func (s *Service) SetTokenByUserId(ctx context.Context, req *authsrvproto.CSTokenSet, rsp *authsrvproto.SCTokenSet) error {
	log.Debug("[SetToken] 收到创建token请求")

	token, err := token.GetTokenService().SetToken(&token.Subject{
		Id: strconv.FormatInt(req.UserId, 10),
	})

	if err != nil {
		rsp.Error = &authsrvproto.Error{
			Detail: err.Error(),
		}

		log.Debugf("[SetToken] token生成失败，err：%s", err)
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
	log.Debug("[ClearToken] 清除用户token")
	err := token.GetTokenService().ClearToken(req.Token)
	if err != nil {
		rsp.Error = &authsrvproto.Error{
			Detail: err.Error(),
		}

		log.Debugf("[ClearToken] 清除用户token失败，err：%s", err)
		return err
	}

	rsp.Error = &authsrvproto.Error{
		Code: 200,
	}

	return nil
}

// 获取缓存的token
func (s *Service) GetTokenByUserId(ctx context.Context, req *authsrvproto.CSTokenGet, rsp *authsrvproto.SCTokenGet) error {
	log.Debug("[GetToken] 获取缓存的token，%d", req.UserId)
	token, err := token.GetTokenService().GetToken(&token.Subject{
		Id: strconv.FormatInt(req.UserId, 10),
	})
	if err != nil {
		rsp.Error = &authsrvproto.Error{
			Detail: err.Error(),
		}

		log.Debugf("[GetToken] 获取缓存的token失败，err：%s", err)
		return err
	}

	rsp.Error = &authsrvproto.Error{
		Code: 200,
	}
	rsp.Token = token

	return nil
}
