package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-plugins/config/source/grpc"
	openTrace "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	log "github.com/sirupsen/logrus"
	"os"
	authsrvcontroller "project/shop/auth_srv/controller"
	authsrvmodel "project/shop/auth_srv/model"
	authsrvproto "project/shop/auth_srv/proto"
	"project/shop/basic"
	basiccommon "project/shop/basic/common"
	"project/shop/basic/config"
	tracer "project/shop/common/tracer/jaeger"
)

var (
	appName = "auth_srv"
	cfg     = &authCfg{}
)

type authCfg struct {
	basiccommon.AppCfg
}

func main() {
	initLog()
	initConfig()

	micReg := etcd.NewRegistry(registryOptions)

	t, io, err := tracer.NewTracer(cfg.Name, "")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 新建服务
	service := micro.NewService(
		micro.Name(cfg.Name),   // 	服务名字
		micro.Registry(micReg), // etcd配置信息
		micro.Version(cfg.Version),
		micro.Address(cfg.Addr()), // 服务地址
		micro.WrapHandler(openTrace.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	// 服务器初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化model
			authsrvmodel.Init()
		}),
	)

	// 注册服务
	authsrvproto.RegisterUserHandler(service.Server(), new(authsrvcontroller.Service))

	log.Info("authsrv: 启动authsrv服务...")

	// 启动服务
	if err := service.Run(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("authsrv:  tcp accept错误")
	}

}

func registryOptions(ops *registry.Options) {
	etcdCfg := &basiccommon.Etcd{}
	err := config.GetConfigurator().App("etcd", etcdCfg)
	if err != nil {
		panic(err)
	}

	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.Host, etcdCfg.Port)}
}

// 初始化配置信息，监听配置变动
func initConfig() (err error) {
	source := grpc.NewSource(
		grpc.WithAddress(basiccommon.EtcdAddr), // 配置地址
		grpc.WithPath("conf"),                  // 对应配置
	)

	basic.Init(config.WithSource(source))

	err = config.GetConfigurator().App(appName, cfg)
	if err != nil {
		log.WithFields(log.Fields{
			"appName": appName,
			"error":   err,
		}).Fatal("authsrv: 初始化配置失败")
		return
	}

	log.WithFields(log.Fields{
		"cfg": *cfg,
	}).Info("authsrv: 配置信息")

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
