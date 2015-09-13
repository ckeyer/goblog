package controllers

type IndexController struct {
	Controller
}

func (c *IndexController) Get() {
	log.Debug("Index....")
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}
