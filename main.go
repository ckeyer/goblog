package main

import (
	_ "blog/routers"
	"encoding/base64"
	"github.com/astaxie/beego"
	"strings"
)

func main() {
	beego.TemplateLeft = "<%"
	beego.TemplateRight = "%>"
	beego.AddFuncMap("DECODEBASE64", decodeBase64)
	beego.Run()
}
func decodeBase64(s string) string {
	s = strings.Replace(s, "+", "-", -1)
	s = strings.Replace(s, "/", "_", -1)
	v, _ := base64.URLEncoding.DecodeString(s)
	return string(v)
}
