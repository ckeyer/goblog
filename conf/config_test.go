package conf

import (
	"testing"
)

// TestConfig 测试配置相关
func TestConfig(t *testing.T)  {
	err:=LoadConf("./v2.json")
	if err != nil {
		t.Errorf("test loadConfig error")
	}
	if GetConf()==nil {
		t.Error("获取配置为nil")
	}
	
}
