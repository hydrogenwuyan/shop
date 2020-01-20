package config

import (
	"fmt"
	"github.com/micro/go-micro/config"
	log "github.com/sirupsen/logrus"
	"sync"
)

var (
	once sync.Once
	c    *configurator // 默认配置器
)

type Configurator interface {
	App(name string, config interface{}) (err error)
}

type configurator struct {
	conf config.Config
}

// 根据app name获取配置
func (c *configurator) App(name string, config interface{}) (err error) {
	v := c.conf.Get(name)
	if v != nil {
		err = v.Scan(config)
		if err != nil {
			log.WithFields(log.Fields{
				"appName": name,
				"error":   err,
			}).Error("basic: 获取配置失败")
			return
		}
	} else {
		err = fmt.Errorf("config get fail")
		log.WithFields(log.Fields{
			"appName": name,
			"error":   err,
		}).Error("basic: 获取配置失败")
		return
	}

	return
}

func (c *configurator) init(ops Options) (err error) {
	once.Do(func() {
		log.Info("basic: configurator开始初始化...")

		c.conf = config.NewConfig()

		// 加载配置
		err := c.conf.Load(ops.Sources...)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("basic: 加载sources失败")
			return
		}

		// 开始监听配置变动
		go func() {
			watch, err := c.conf.Watch()
			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
				}).Error("basic: 监听配置变化失败")
				return
			}

			// 开始监听
			for {
				// TODO: 配置请求频率过高,待优化（现在是1s一次）
				v, err := watch.Next()
				if err != nil {
					log.WithFields(log.Fields{
						"error": err,
					}).Error("basic: 监听配置变动失败")
					return
				}

				log.WithFields(log.Fields{
					"v": string(v.Bytes()),
				}).Debug("basic: 配置变动")
			}
		}()
	})

	return
}

// 初始化配置
func Init(opts ...Option) {

	ops := Options{}
	for _, o := range opts {
		o(&ops)
	}

	c = &configurator{}
	c.init(ops)
}

// 获取配置器
func GetConfigurator() Configurator {
	return c
}
