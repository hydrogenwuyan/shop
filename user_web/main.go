package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/config/source/grpc"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"project/shop/basic"
	basiccommon "project/shop/basic/common"
	"project/shop/basic/config"
	"project/shop/user_web/controller"
	"time"
)

var (
	appName = "user_web"
	cfg     = &userCfg{}
)

type userCfg struct {
	basiccommon.AppCfg
}

func main() {
	initLog()
	initConfig()

	etcdReg := etcd.NewRegistry(registryOptions)

	// 新建服务
	service := web.NewService(
		web.Name(cfg.Name),
		web.Version(cfg.Version),
		web.RegisterTTL(time.Second*15), // 数据包生存时间
		web.RegisterInterval(time.Second*10),
		web.Registry(etcdReg),
		web.Address(cfg.Addr()),
	)

	// 初始化服务
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				// 初始化handler
				controller.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}

	// 注册路由
	service.Handle("/user/login", http.HandlerFunc(controller.Login))
	service.Handle("/user/logout", http.HandlerFunc(controller.Logout))

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

// 读取etcd配置信息
func registryOptions(ops *registry.Options) {
	etcdCfg := &basiccommon.Etcd{}
	err := config.GetConfigurator().App("etcd", etcdCfg)
	if err != nil {
		panic(err)
	}

	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.Host, etcdCfg.Port)}
}

// 读取配置信息
func initConfig() (err error) {
	source := grpc.NewSource(
		grpc.WithAddress(basiccommon.EtcdAddr),
		grpc.WithPath("conf"),
	)

	basic.Init(config.WithSource(source))

	err = config.GetConfigurator().App(appName, cfg)
	if err != nil {
		log.WithFields(log.Fields{
			"appName": appName,
			"error":   err,
		}).Fatal("userweb: 初始化配置失败")
		return
	}

	log.WithFields(log.Fields{
		"cfg": *cfg,
	}).Info("userweb: 配置信息")

	return
}

func initLog() {
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}
