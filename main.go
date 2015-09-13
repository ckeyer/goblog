package main

import (
	_ "github.com/ckeyer/goblog/routers"
	"github.com/astaxie/beego"
	"github.com/ckeyer/goblog/conf"
)

func main() {
	conf.LoadConf("conf/v2.json")
	beego.Run()
}

