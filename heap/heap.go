package heap

import "sync"

// Heap 用来存储每个Node在被NodeRunner运行时所使用的全局变量，命名为堆
type Heap struct {
	Kv     map[string]interface{} // 堆的key-value对
	Status int                    // 状态
	mutex  *sync.Mutex            //加锁
}

// Update 从其它heap中更新
func (H *Heap) Update(h Heap) error {
	for key, value := range h.Kv {
		H.Kv[key] = value
	}
	return nil
}

// Merge 合并其它heap
func (H *Heap) Merge(heapList ...Heap) error {
	result := make(map[string]interface{})
	for key, value := range H.Kv {
		temValueList := make([]interface{}, 0, len(heapList))
		temValueList = append(temValueList, value)
		for _, heapEle := range heapList {
			if heapEleValue, OK := heapEle.Kv[key]; OK {
				temValueList = append(temValueList, heapEleValue)
			}
		}
		if len(temValueList) == 1 {
			continue
		}
		mergedValue, err := specialEle(temValueList...)
		if err != nil {
			return err
		}
		result[key] = mergedValue
	}
	H.Kv = result
	return nil
}

//注意，入参不能为slice、map、function类型,获取出现次数最小的结果
func specialEle(interfaceList ...interface{}) (interface{}, error) {
	countMap := make(map[interface{}]int, 0)
	for _, v := range interfaceList {
		countMap[v]++
	}
	var currentMinEle interface{}
	currentMinIndex := 99999
	for k, index := range countMap {
		if index == 1 {
			return k, nil
		}
		if index < currentMinIndex {
			currentMinEle = k
			currentMinIndex = index
		}
	}
	return currentMinEle, nil
}
