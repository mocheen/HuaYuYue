package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/resolver"
	"net"
	"time"
	"user/conf"
	"user/discovery"
	"user/internal/handler"
	"user/internal/repository"
	"user/internal/repository/query"
	service "user/internal/service/pb"
	"user/internal/service/role"
)

func main() {
	conf.GetConf()
	repository.InitDB()
	query.SetDefault(repository.DB)
	handler.Init()

	//etcd
	etcdAddress := []string{viper.GetString("etcd.address")}
	// 注册
	etcdRegister := discovery.NewRegister(etcdAddress, logrus.New())
	grpcAddress := viper.GetString("server.grpcAddress")
	defer etcdRegister.Stop()
	userNode := discovery.Server{
		Name:    viper.GetString("server.domain"),
		Addr:    grpcAddress,
		Version: viper.GetString("server.version"),
	}

	srv := grpc.NewServer()
	defer srv.Stop()

	service.RegisterUserServiceServer(srv, handler.GetUserSrv())

	// 发现
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	resolver.Register(discovery.NewResolver(etcdAddress, logger))

	// 创建 gRPC 连接并生成客户端实例
	roleServiceConn := getConn("role", etcdAddress, logger)
	handler.RoleServiceClient = role.NewRoleServiceClient(roleServiceConn)

	defer roleServiceConn.Close()

	// 监听
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}
	if _, err := etcdRegister.Register(userNode, 30); err != nil {
		panic(err)
	}
	if err = srv.Serve(lis); err != nil {
		panic(err)
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
