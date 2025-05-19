package main

import (
	"gateway/conf"
	"gateway/discovery"
	"gateway/idl/role"
	"gateway/idl/user"
	"gateway/internal/handler"
	"gateway/routes"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/resolver"
	"time"
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
	userServiceConn := getConn(conf.Conf.Services["user"].Name, etcdAddress, logger)
	roleServiceConn := getConn(conf.Conf.Services["role"].Name, etcdAddress, logger)

	defer userServiceConn.Close()
	defer roleServiceConn.Close()

	// 创建 Client 实例
	userServiceClient := user.NewUserServiceClient(userServiceConn)

	roleServiceClient := role.NewRoleServiceClient(roleServiceConn)

	// 传递客户端实例而非连接对象
	ginRouter := routes.NewRouter(userServiceClient, roleServiceClient)
	if err := ginRouter.Run(":4000"); err != nil {
		logger.Fatalf("failed to start gateway: %v", err)
	}
}

func getConn(serviceName string, etcdAddress []string, logger *logrus.Logger) *grpc.ClientConn {
	Conn, err := grpc.Dial(
		"discovery:///"+serviceName,
		grpc.WithDefaultServiceConfig(`{
        "loadBalancingPolicy": "round_robin",
        "healthCheckConfig": {
            "serviceName": "user",
            "timeout": "5s"
        }
    }`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                20 * time.Second, // 每 20 秒发送 PING
			Timeout:             5 * time.Second,  // 等待 ACK 超时
			PermitWithoutStream: true,             // 无活跃流时也保活
		}),
		grpc.WithResolvers(discovery.NewResolver(etcdAddress, logger)),
	)
	if err != nil {
		panic(err)
	}
	return Conn
}
