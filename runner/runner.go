package runner

import "garten/node"

// NodeRunner 流水线上的一号员工

type NodeRunner struct {
	Probe  *node.Node             //当前运行的节点
	Heap   map[string]interface{} //全局变量，堆的概念
	Status uint                   //当前ChainNode的状态 ，为了方便扩展使用uint类型
}

var NodeRunnerList []NodeRunner

func run(wait bool) (map[string]interface{}, error) {

}
