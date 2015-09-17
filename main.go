package main

import (
	"github.com/astaxie/beego"
	"github.com/ckeyer/goblog/conf"
	"github.com/ckeyer/goblog/libs"
	"github.com/ckeyer/goblog/modules"
	"github.com/ckeyer/goblog/routers"
)

var log = libs.GetLogger()

func init() {
	conf.LoadConf("conf/v2.json")

	err := modules.LoadBlogs(conf.GetConf().BlogDir)
	if err != nil {
		log.Error(err)
	}
}

func main() {
	BeegoInit()
	beego.Run()
}

func BeegoInit() {
	beego.TemplateLeft = "<<<"
	beego.TemplateRight = ">>>"
	routers.LoadRouters()
}
