package node

// Node 流水线上的一个产品
type Node struct {
	RawData    interface{}            //要运行的数据
	Next       *Node                  // 接下来要运行的内容
	Stack      map[string]interface{} // 当前节点内生效的变量
	Status     uint                   //当前Node的状态 ，为了方便扩展使用uint类型
	RunnerData interface{}
}
