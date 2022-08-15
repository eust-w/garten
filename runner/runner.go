package runner

import (
	"fmt"
	"garten/heap"
	"garten/node"
)

// NodeRunner 流水线上的一号员工

//type NodeRunner struct {
//	Probe  *node.Node             //当前运行的节点
//	Heap   map[string]interface{} //全局变量，堆的概念
//	Status uint                   //当前ChainNode的状态 ，为了方便扩展使用uint类型
//}

func NodeRunner(nodeIn node.Node, heapIn heap.Heap) (heap.Heap, error) {
	env := heapIn.Kv

}

func parserNode(nodeIn node.Node, heapIn heap.Heap) {
	nodeTag := nodeIn.Tag
	if nodeTag == "parallel" {
		parallelMap := nodeIn.RawData.(map[string]interface{})
		resultHeapList := make([]heap.Heap,0,len(parallelMap))
		for k, v := range parallelMap {
			newNode := node.Node{RawData: v, Tag: k}
			newHeap := heapIn
			go func(nodeIn node.Node, heapIn heap.Heap) {
				result, err := NodeRunner(nodeIn,heapIn)
				if err != nil{
					fmt.Println(" NodeRunner err ")
				}
				resultHeapList = append(resultHeapList, result)
			}(newNode, newHeap)
		}
	}
	heapOut,flag,err := baseParser(nodeTag,nodeIn,heapIn){
		if flag{
			return
		}
	}
}

