package main

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
	proto "github.com/micro/go-plugins/config/source/grpc/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"project/shop/config_srv/service"
)

var (
	apps = []string{"conf"}
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal("configsrv: 服务崩溃")
		}
	}()

	initLog()

	err := loadAndWatchConfigFile()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("configsrv: loadAndWatchConfigFile 失败")
		return
	}

	// 新建grpc服务
	grpcServer := grpc.NewServer()
	proto.RegisterSourceServer(grpcServer, new(service.ConfigService))

	// 创建tcp链接
	l, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("configsrv: tcp listen错误")
	}

	log.Info("configsrv: 启动configsrv服务...")

	// 启动
	err = grpcServer.Serve(l)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("configsrv: tcp accept错误")
	}
}

// 加载并监听配置文件
func loadAndWatchConfigFile() (err error) {
	for _, app := range apps {
		err = config.Load(file.NewSource(file.WithPath("./conf/" + app + ".yml")))
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Fatal("configsrv: 加载config文件失败")
			return
		}
	}

	watch, err := config.Watch()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("configsrv: 创建watcher失败")
		return
	}

	// 启动监听
	go func() {
		for {
			v, err := watch.Next()
			if err != nil {
				log.WithFields(log.Fields{
					"v":     v,
					"error": err,
				}).Error("configsrv: 监听配置变化错误")
				return
			}

			log.WithFields(log.Fields{
				"v":     string(v.Bytes()),
				"error": err,
			}).Debug("configsrv: 配置信息改变")
		}
	}()

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
