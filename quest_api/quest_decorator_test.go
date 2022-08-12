package quest_api

import (
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
