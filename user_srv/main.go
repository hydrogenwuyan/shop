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
	"project/shop/basic"
	basiccommon "project/shop/basic/common"
	"project/shop/basic/config"
	tracer "project/shop/common/tracer/jaeger"
	usersrvcontroller "project/shop/user_srv/controller"
	usersrvmodel "project/shop/user_srv/model"
	usersrvproto "project/shop/user_srv/proto"
	"time"
)

var (
	appName = "user_srv"
	cfg     = &userCfg{}
)

type userCfg struct {
	basiccommon.AppCfg
}

func main() {
	initLog()
	initConfig()

	etcdReg := etcd.NewRegistry(registryOptions)

	t, io, err := tracer.NewTracer(cfg.Name, "")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 新建服务
	service := micro.NewService(
		micro.Name(cfg.Name),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		micro.Registry(etcdReg),
		micro.Version("latest"),
		micro.WrapHandler(openTrace.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化model层
			usersrvmodel.Init()
		}),
	)

	// 注册服务
	usersrvproto.RegisterUserHandler(service.Server(), new(usersrvcontroller.Service))

	// 启动服务
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
		}).Fatal("usersrv: 初始化配置失败")
		return
	}

	log.WithFields(log.Fields{
		"cfg": *cfg,
	}).Info("usersrv: 配置信息")

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
