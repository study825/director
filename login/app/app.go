package app

import (
	"github.com/dobyte/due/v2/cluster/node"
	"github.com/study825/director/login/logic"
)

func Init(proxy *node.Proxy) {
	logic.NewCore(proxy).Init()
}
