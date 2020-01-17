package main

import (
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
	"github.com/micro/go-micro/util/log"
	proto "github.com/micro/go-plugins/config/source/grpc/proto"
	"google.golang.org/grpc"
	"net"
	"project/shop/config_srv/service"
)

var (
	apps = []string{"conf"}
)

func main() {
	defer func() {
		if r := recover();r != nil {
			log.Logf("config main panic, %v", r)
		}
	}()

	err := loadAndWatchConfigFile()
	if err != nil {
		log.Fatal("loadAndWatchConfigFile fail")
		return
	}

	// 新建grpc服务
	grpcServer := grpc.NewServer()
	proto.RegisterSourceServer(grpcServer, new(service.ConfigService))

	// 创建tcp链接
	l, err := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil {
		log.Fatalf("listen fail, err: %v", err)
	}

	log.Info("start config grpc server")

	// 启动
	err = grpcServer.Serve(l)
	if err != nil {
		log.Fatalf("tcp change grpc fail, err: %v", err)
	}
}

// 加载并监听配置文件
func loadAndWatchConfigFile() (err error){
	for _, app := range apps {
		err = config.Load(file.NewSource(file.WithPath("./conf/"+app+".yml")))
		if err != nil {
			log.Fatalf("load config fail, err: %v", err)
			return
		}
	}

	watch, err := config.Watch()
	if err != nil {
		log.Fatalf("watch config fail, err: %v", err)
		return
	}

	// 启动监听
	go func() {
		for {
			v, err := watch.Next()
			if err != nil {
				log.Errorf("watch next config fail, err %v", err)
				return
			}

			log.Debugf("watch config changed, %s", string(v.Bytes()))
		}
	}()

	return
}
