package runner

import (
	"garten/heap"
	"garten/node"
)

func baseParser(tag string, nodeIn node.Node, heapIn heap.Heap) (heap.Heap, bool, error) {
	switch tag {
	case "ssh":

	}
}

func ssh(nodeIn node.Node, heapIn heap.Heap) {
	env := nodeIn.RawData.(map[string]string)
	var cmd string
	if v,OK := env["cmd"];OK{
		cmd = v
	}else {
		//cmd必须要有
		return
	}

	if v,OK := env["cluster"];OK{
		sshCluster(nodeIn,heapIn)
		return
	}
	if v,OK := env[]

}

func sshCluster(nodeIn node.Node, heapIn heap.Heap) {}