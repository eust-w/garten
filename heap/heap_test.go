package heap

import (
	"fmt"
	"testing"
)

func TestHeap_specialEle(t *testing.T) {
	p0 := []string{"1"}
	p1 := []string{"1"}
	//p0 := 3.23
	//p1 := 3.23
	p2 := 2
	p3 := "13"
	p4 := 3
	p5 := "13"
	p6 := 2
	p7 := "13"
	out, err := specialEle(p0, p1, p2, p3, p4, p5, p6, p7)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
