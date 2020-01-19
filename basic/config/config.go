package config

import (
	"fmt"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
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

func (c *configurator) App(name string, config interface{}) (err error) {
	v := c.conf.Get(name)
	if v != nil {
		err = v.Scan(config)
		if err != nil {
			log.Errorf("config scan fail, name:%s err: %v", name, err)
			return
		}
	} else {
		err = fmt.Errorf("config get fail")
		log.Error(err)
		return
	}

	return
}

func (c *configurator) init(ops Options) (err error) {
	once.Do(func() {
		log.Info("start init configurator")

		c.conf = config.NewConfig()

		// 加载配置
		err := c.conf.Load(ops.Sources...)
		if err != nil {
			log.Errorf("load source fail, err: %v", err)
			return
		}

		// 开始监听配置变动
		go func() {
			watch, err := c.conf.Watch()
			if err != nil {
				log.Errorf("watch fail, err: %v", err)
				return
			}

			// 开始监听
			for {
				v, err := watch.Next()
				if err != nil {
					log.Errorf("watch next, err: %v", err)
					return
				}

				log.Debugf("config changed, %s", string(v.Bytes()))
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
