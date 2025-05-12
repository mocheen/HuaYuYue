package main

import (
	"gateway/conf"
	"gateway/discovery"
	"gateway/idl/user"
	"gateway/internal/handler"
	"gateway/routes"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

func main() {
	conf.GetConf()
	handler.Init()

	// 服务发现初始化
	etcdAddress := []string{viper.GetString("etcd.address")}
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	resolver.Register(discovery.NewResolver(etcdAddress, logger))

	// 创建 gRPC 连接并生成客户端实例
	userServiceConn, err := grpc.Dial(
		"discovery:///user",
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	defer userServiceConn.Close()

	// 关键修改：创建 UserServiceClient 实例
	userServiceClient := user.NewUserServiceClient(userServiceConn)

	// 传递客户端实例而非连接对象
	ginRouter := routes.NewRouter(userServiceClient)
	if err := ginRouter.Run(":4000"); err != nil {
		logger.Fatalf("failed to start gateway: %v", err)
	}
}
