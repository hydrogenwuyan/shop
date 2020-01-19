package token

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"project/shop/basic/config"
	"project/shop/common"
	"sync"
)

var (
	once        sync.Once
	s           *tokenService
	redisClient *redis.Client
	cfg         *common.Jwt
)

type TokenService interface {
	// 设置token
	SetToken(subject *Subject) (token string, err error)
	// 清理token
	ClearToken(token string) (err error)
	// 获取token
	GetToken(subject *Subject) (token string, err error)
}

func Init() (err error) {
	once.Do(func() {
		log.Info("token服务开始初始化...")

		cfg = &common.Jwt{}
		err := config.GetConfigurator().App("jwt", cfg)
		if err != nil {
			panic(err)
		}

		redisClient = common.Redis()
		s = &tokenService{}
	})

	return
}

func GetTokenService() TokenService {
	return s
}
