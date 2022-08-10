package garten

// Node 流水线上的一个产品
type Node struct {
	RawData    interface{}            //要运行的数据
	Next       *Node                  // 接下来要运行的内容
	Stack      map[string]interface{} // 当前节点内生效的变量
	Status     uint                   //当前Node的状态 ，为了方便扩展使用uint类型
	RunnerData interface{}
}

// NodeRunner 流水线上的一号员工
type NodeRunner struct {
	Probe  *Node                  //当前运行的节点
	Heap   map[string]interface{} //全局变量，堆的概念
	Status uint                   //当前ChainNode的状态 ，为了方便扩展使用uint类型
}
