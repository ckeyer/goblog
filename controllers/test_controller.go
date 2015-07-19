package controllers

import (
// _ "container/list"
// "github.com/vmihailenco/redis"
// "github.com/ckeyer/goblog/models"
)

type TestController struct {
	BaseController
}

func (this *TestController) Get() {
	this.Data["PageTitle"] = "测试"
	// this.Data["TestStr"] = "s"
	this.Data["CusScripts"] = `<script src="` + STATIC_URL_JS + `test.js"></script>`
	this.TplNames = "test.tpl"
}
func (t *TestController) Post() {
	t.Ctx.WriteString("YES")
}
func Test() {
	log.Debug("######## TEST ########")

	log.Debug("######## TEST ########")
}
