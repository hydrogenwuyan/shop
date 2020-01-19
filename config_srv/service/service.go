package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
	proto "github.com/micro/go-plugins/config/source/grpc/proto"
	"strings"
	"time"
)

// 配置服务
type ConfigService struct{}

// 读取配置
func (s ConfigService) Read(ctx context.Context, req *proto.ReadRequest) (resp *proto.ReadResponse, err error) {
	appName := parsePath(req.Path)
	fmt.Println("*********appName********", appName)
	resp = &proto.ReadResponse{
		ChangeSet: getConfig(appName),
	}
	return
}

// 监听配置变动，发送变动的配置给各个服务
func (s ConfigService) Watch(req *proto.WatchRequest, server proto.Source_WatchServer) (err error) {
	appName := parsePath(req.Path)
	resp := &proto.WatchResponse{
		ChangeSet: getConfig(appName),
	}

	err = server.Send(resp)
	if err != nil {
		log.Errorf("source watch server send fail, err: %v", err)
	}

	return
}

// 获取配置
func getConfig(appName string) *proto.ChangeSet {
	bytes := config.Get(appName).Bytes()
	return &proto.ChangeSet{
		Data:      bytes,
		Checksum:  fmt.Sprintf("%x", md5.Sum(bytes)),
		Format:    "yml",
		Source:    "file",
		Timestamp: time.Now().Unix(),
	}
}

// 解析path
func parsePath(path string) (appName string) {
	paths := strings.Split(path, "/")

	if paths[0] == "" && len(paths) > 1 {
		return paths[1]
	}

	return paths[0]
}
