package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"project/shop/basic"
	"project/shop/basic/config"
	"sync"
)

var (
	once        sync.Once
	redisClient *redis.Client
)

// redis 配置
type redisConfig struct {
	Enabled  bool   `json:"enabled"`
	Address  string `json:"address"`
	Port     int32  `json:"port"`
	Password string `json:"passWord"`
	DBNum    int    `json:"dbNum"`
	Timeout  int    `json:"timeout"`
	//Sentinel *RedisSentinel `json:"sentinel"`
}

func initRedis() {
	once.Do(func() {
		cfg := &redisConfig{}
		c := config.GetConfigurator()
		err := c.App("redis", cfg)
		if err != nil {
			log.Fatalf("common: 获取redis配置失败, error: %v", err)
			return
		}

		if !cfg.Enabled {
			log.Info("未启用redis")
			return
		}

		redisClient = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", cfg.Address, cfg.Port),
			Password: cfg.Password,
			DB:       cfg.DBNum, // use default DB
		})
	})

	return
}

func init() {
	basic.Register(initRedis)
}

// 获取redis
func Redis() *redis.Client {
	return redisClient
}
