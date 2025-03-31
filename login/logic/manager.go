package logic

import (
	"github.com/dobyte/due/v2/cluster/node"
	"golang.org/x/sync/singleflight"
	"sync"
	"sync/atomic"
)

type Manager struct {
	proxy *node.Proxy        // 代理API
	id    atomic.Int64       // 聊天室自增ID
	sfg   singleflight.Group // single flight
	users sync.Map           // 用户列表
	rooms sync.Map           // 聊天室列表
}

func newManager(proxy *node.Proxy) *Manager {
	return &Manager{proxy: proxy}
}
