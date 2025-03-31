package logic

import (
	"github.com/dobyte/due/v2/cluster/node"
)

type Core struct {
	proxy   *node.Proxy
	manager *Manager
}

func NewCore(proxy *node.Proxy) *Core {
	return &Core{proxy: proxy, manager: newManager(proxy)}
}

func (c *Core) Init() {
	c.proxy.Router().Group(func(group *node.RouterGroup) {

	})
}
