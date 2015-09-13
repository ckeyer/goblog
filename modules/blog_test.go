package modules

import (
	"testing"
)

// TestNewBlog ...
func TestNewBlog(t *testing.T)  {
	b,err:=NewBlog("../blog/study-golang1.md")
	if err != nil {
		t.Error( "加载Blog错误", err)
	}
	_=b

}
