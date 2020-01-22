package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/micro/go-micro/config"
	proto "github.com/micro/go-plugins/config/source/grpc/proto"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

// 配置服务
type ConfigService struct{}

// 读取配置
func (s ConfigService) Read(ctx context.Context, req *proto.ReadRequest) (resp *proto.ReadResponse, err error) {
	log.WithFields(log.Fields{
		"ReadRequest": *req,
	}).Debug("configsrv: 读取配置")

	appName := parsePath(req.Path)
	resp = &proto.ReadResponse{
		ChangeSet: getConfig(appName),
	}
	return
}

// 监听配置变动，发送变动的配置给各个服务
func (s ConfigService) Watch(req *proto.WatchRequest, server proto.Source_WatchServer) (err error) {
	//log.WithFields(log.Fields{
	//	"watchRequest": *req,
	//}).Debug("configsrv: 请求配置信息")

	appName := parsePath(req.Path)
	resp := &proto.WatchResponse{
		ChangeSet: getConfig(appName),
	}

	err = server.Send(resp)
	if err != nil {
		log.WithFields(log.Fields{
			"watchResponse": *resp,
			"error":         err,
		}).Error("configsrv: 发送配置信息失败")
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
