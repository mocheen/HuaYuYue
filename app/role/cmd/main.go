package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"role/conf"
	"role/discovery"
	"role/internal/handler"
	"role/internal/repository"
	"role/internal/repository/query"
	service "role/internal/service/pb"
)

func main() {
	conf.GetConf()
	repository.InitDB()
	query.SetDefault(repository.DB)

	//etcd
	etcdAddress := []string{viper.GetString("etcd.address")}
	etcdRegister := discovery.NewRegister(etcdAddress, logrus.New())
	grpcAddress := viper.GetString("server.grpcAddress")
	defer etcdRegister.Stop()
	roleNode := discovery.Server{
		Name:    viper.GetString("server.domain"),
		Addr:    grpcAddress,
		Version: viper.GetString("server.version"),
	}

	srv := grpc.NewServer()
	defer srv.Stop()

	service.RegisterRoleServiceServer(srv, handler.GetRoleSrv())

	// 监听
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}
	if _, err := etcdRegister.Register(roleNode, 30); err != nil {
		panic(err)
	}
	if err = srv.Serve(lis); err != nil {
		panic(err)
	}

}
