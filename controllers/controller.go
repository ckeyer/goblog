package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ckeyer/goblog/conf"
	"github.com/ckeyer/goblog/libs"
	"strings"
)

var (
	website *conf.WebSite
	log     = libs.GetLogger()
)

type Controller struct {
	beego.Controller
}

// 是否是通过允许的域名访问
func (c *Controller) IsAllowHost() bool {
	host := c.Ctx.Input.Host()
	for _, v := range website.EnableDomain {
		if strings.Index(host, v) >= 0 {
			return true
		}
	}
	log.Debugf("Not Enable Domain %s\n", c.Ctx.Input.Host())
	return false
}

//  (b *Controller)SetPateTitle 设置页面显示标题
func (c *Controller) SetPageTitle(title string) {
	c.Data["PageTitle"] = title
}

// (b *Controller)SetKeyWord 设置或更新Meta关键字
func (c *Controller) SetKeyWord(args ...string) {
	c.Data["Keywords"] = strings.Join(args, ",")
}

// (b *Controller)AddKeyWord 添加Meta关键字
func (c *Controller) AddKeyWord(args ...string) {
	if c.Data["Keywords"] != nil {
		keyword := strings.Split(fmt.Sprint(c.Data["Keywords"]), ",")
		c.Data["Keywords"] = strings.Join(append(keyword, args...), ",")
	} else {
		c.SetKeyWord(args...)
	}
}

// (b *Controller)SetDescript 设置Mets描述
func (c *Controller) SetDescript(des string) {
	c.Data["Description"] = des
}

// (b *Controller)AddJsScript 按配置的js路径添加js文件
func (c *Controller) AddJsScript(args ...string) {
	c.AddCustomJsScript(website.JsUrl, args...)
}

// (b *Controller)AddCustomJsScript 添加自定义js
func (c *Controller) AddCustomJsScript(src_url string, args ...string) {
	var jstags []string

	if c.Data["Scripts"] != nil {
		jstags = strings.Split(fmt.Sprint(c.Data["Scripts"]), "\n")
	}
	for _, js := range args {
		newtag := fmt.Sprintf(`<script type="text/javascript" src="%s%s"></script>`, src_url, js)
		jstags = append(jstags, newtag)
	}
	c.Data["Scripts"] = strings.Join(jstags, "\n")

}

// (b *Controller)AddCssStyle 按配置的css路径添加css文件
func (c *Controller) AddCssStyle(args ...string) {
	c.AddCustomCssStyle(website.CssUrl, args...)
}

// (b *Controller)AddCustomCssStyle 添加自定义css
func (c *Controller) AddCustomCssStyle(src_url string, args ...string) {
	var csstags []string

	if c.Data["Styles"] != nil {
		csstags = strings.Split(fmt.Sprint(c.Data["Styles"]), "\n")
	}
	for _, css := range args {
		newtag := fmt.Sprintf(`<link rel="stylesheet" media="screen" type="text/css" href="%s%s"/>`, src_url, css)
		csstags = append(csstags, newtag)
	}
	c.Data["Styles"] = strings.Join(csstags, "\n")
}
