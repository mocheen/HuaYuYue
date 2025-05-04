package discovery

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
)

type Resolver struct {
	schema      string
	EtcdAddrs   []string
	DialTimeout int

	closeCh      chan struct{}
	watchCh      clientv3.WatchChan
	cli          *clientv3.Client
	keyPrifix    string
	srvAddrsList []resolver.Address

	cc     resolver.ClientConn
	logger *logrus.Logger
}

func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r.cc = cc
	serviceName := strings.TrimPrefix(target.Endpoint(), "/")
	r.keyPrifix = fmt.Sprintf("/%s/%s/", r.schema, serviceName) // 格式: /etcd/user/

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   r.EtcdAddrs,
		DialTimeout: time.Duration(r.DialTimeout) * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("etcd client failed: %v", err)
	}
	r.cli = cli
	r.closeCh = make(chan struct{})

	go r.watch()
	r.resolve()

	return r, nil
}

func (r *Resolver) watch() {
	defer r.cli.Close()

	r.watchCh = r.cli.Watch(context.Background(), r.keyPrifix, clientv3.WithPrefix())
	for {
		select {
		case <-r.closeCh:
			return
		case resp, ok := <-r.watchCh:
			if !ok {
				r.logger.Warn("watch channel closed")
				return
			}
			for range resp.Events {
				r.resolve() // 监听到变化时触发更新
			}
		}
	}
}

func (r *Resolver) resolve() {
	resp, err := r.cli.Get(context.Background(), r.keyPrifix, clientv3.WithPrefix())
	if err != nil {
		r.logger.Errorf("etcd get failed: %v", err)
		return
	}

	var addrs []resolver.Address
	for _, kv := range resp.Kvs {
		server, err := ParseValue(kv.Value)
		if err != nil {
			r.logger.Warnf("parse etcd value failed: %v", err)
			continue
		}
		addr := resolver.Address{Addr: server.Addr}
		if !Exist(addrs, addr) {
			addrs = append(addrs, addr)
		}
	}

	if len(addrs) > 0 {
		r.cc.UpdateState(resolver.State{Addresses: addrs})
		r.logger.Infof("updated addresses: %v", addrs)
	}
}

func (r *Resolver) Scheme() string                        { return r.schema }
func (r *Resolver) Close()                                { close(r.closeCh) }
func (r *Resolver) ResolveNow(resolver.ResolveNowOptions) {}

func NewResolver(etcdAddress []string, logger *logrus.Logger) *Resolver {
	return &Resolver{
		schema:       "etcd",
		EtcdAddrs:    etcdAddress,
		DialTimeout:  3,
		closeCh:      make(chan struct{}),         // 用于优雅关闭 watch 协程
		srvAddrsList: make([]resolver.Address, 0), // 初始化服务地址列表
		logger:       logger,
	}
}
