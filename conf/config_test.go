package conf

import (
	"testing"
)

// TestConfig conf包测试
func TestConfig(t *testing.T) {
	SetFilePath("./config.json")
	c, err := GetConfig()
	if err != nil {
		t.Error(err.Error())
	}

	if c == nil {
		t.Error("config is nil")
		return
	}
	if c.App.Port < 80 {
		t.Error("Config Load Error")
	}
	if c.Mysql.GetConnStr() == "" {
		t.Error("Error Mysql COnnstr")
	}
}
