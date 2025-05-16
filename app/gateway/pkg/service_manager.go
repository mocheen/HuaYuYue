package pkg

import (
	"gateway/discovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"sync"
)

type ServiceManager struct {
	clients         sync.Map // 并发安全的存储
	resolverBuilder *discovery.Resolver
}

func NewServiceManager(etcdAddrs []string) *ServiceManager {
	return &ServiceManager{
		resolverBuilder: discovery.NewResolver(etcdAddrs, logrus.New()),
	}
}

func (m *ServiceManager) GetClient(serviceName string, newClientFunc func(cc grpc.ClientConnInterface) interface{}) (interface{}, error) {
	// 检查是否已缓存
	if client, ok := m.clients.Load(serviceName); ok {
		return client, nil
	}

	// 创建新连接
	conn, err := grpc.Dial(
		"discovery:///"+serviceName,
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	// 创建客户端并缓存
	client := newClientFunc(conn)
	m.clients.Store(serviceName, client)

	return client, nil
}
