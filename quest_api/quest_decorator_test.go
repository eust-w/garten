package quest_api

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	para := make(map[string]string)
	para["ip"] = "172.20.1.34"
	para["password"] = "password"
	para["port"] = "8080"
	para["admin"] = "admin"
	c, _ := NewApi(para)
	para2 := make(map[string]interface{})
	para2["zone_name"] = "longtao2"
	para2["description"] = "test"
	Decorator(*c, "CreateZone", para2)
}

func TestDecoratorQueryVmInstance(t *testing.T) {
	para := make(map[string]string)
	para["ip"] = "172.31.10.210"
	para["password"] = "password"
	para["port"] = "8080"
	para["admin"] = "admin"
	c, _ := NewApi(para)
	para2 := make(map[string]interface{})
	para2["uuid"] = "774965dc31794190a8383869c97263d5"
	out, err := Decorator(*c, "QueryVmInstance", para2)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

func TestDecoratorStopVmInstance(t *testing.T) {
	para := make(map[string]string)
	para["ip"] = "172.31.10.210"
	para["password"] = "password"
	para["port"] = "8080"
	para["admin"] = "admin"
	c, _ := NewApi(para)
	para2 := make(map[string]interface{})
	para2["uuid"] = "774965dc31794190a8383869c97263d5"
	out, err := Decorator(*c, "StopVmInstance", para2)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
func TestDecoratorTags(t *testing.T) {
	para := make(map[string]string)
	para["ip"] = "172.31.10.210"
	para["password"] = "password"
	para["port"] = "8080"
	para["admin"] = "admin"
	c, _ := NewApi(para)
	para2 := make(map[string]interface{})
	para2["uuid"] = "774965dc31794190a8383869c97263d5"
	para2["tagUuid"] = "4a9bcad674704c198421dfb41abfaefe"
	//out, err := Decorator(*c, "QueryTag", para2)
	out, err := Decorator(*c, "AttachTagToResources", para2)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

func TestDecoratorTags2(t *testing.T) {
	para := make(map[string]string)
	para["ip"] = "172.31.10.210"
	para["password"] = "password"
	para["port"] = "8080"
	para["admin"] = "admin"
	c, _ := NewApi(para)
	para2 := make(map[string]interface{})
	para2["uuid"] = "774965dc31794190a8383869c97263d5"
	//out, err := Decorator(*c, "QueryTag", para2)
	out, err := Decorator(*c, "ZqlVmTagList", para2)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
