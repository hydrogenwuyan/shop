package common

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/micro/go-micro/util/log"
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
	Port     string `json:"port"`
	Password string `json:"passWord"`
	DBNum    int    `json:"dbNum"`
	Timeout  int    `json:"timeout"`
	//Sentinel *RedisSentinel `json:"sentinel"`
}

func RedisInit() (err error) {
	once.Do(func() {
		cfg := &redisConfig{}
		c := config.GetConfigurator()
		c.App("redis", cfg)

		if !cfg.Enabled {
			log.Logf("未启用redis")
			return
		}

		redisClient = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", cfg.Address, cfg.Port),
			Password: cfg.Password,
			DB:       cfg.DBNum, // use default DB
		})
	})

	return
}

// 获取redis
func Redis() *redis.Client {
	return redisClient
}
