package main

import (
	"encoding/base64"
	"github.com/astaxie/beego"
	_ "github.com/ckeyer/goblog/routers"
	"strings"
)

func main() {
	config_init()
	beego.Run()
}
func config_init() {
	beego.TemplateLeft = "<%"
	beego.TemplateRight = "%>"

	beego.AddFuncMap("STATIC_URL", func() string {
		return beego.AppConfig.String("static_rul")
	})
	beego.AddFuncMap("SITE_URL", func() string {
		return beego.AppConfig.String("site_url")
	})
	beego.AddFuncMap("STATIC_URL_JS", func() string {
		return beego.AppConfig.String("static_url_js")
	})
	beego.AddFuncMap("STATIC_URL_CSS", func() string {
		return beego.AppConfig.String("static_url_css")
	})
	beego.AddFuncMap("STATIC_URL_IMG", func() string {
		return beego.AppConfig.String("static_url_img")
	})
	beego.AddFuncMap("CUSTOM_URL_JS", func() string {
		return beego.AppConfig.String("custom_url_js")
	})
	beego.AddFuncMap("CUSTOM_URL_CSS", func() string {
		return beego.AppConfig.String("custom_url_css")
	})
	beego.AddFuncMap("CUSTOM_URL_IMG", func() string {
		return beego.AppConfig.String("custom_url_img")
	})

	beego.AddFuncMap("DECODEBASE64", func(s string) string {
		s = strings.Replace(s, "+", "-", -1)
		s = strings.Replace(s, "/", "_", -1)
		v, _ := base64.URLEncoding.DecodeString(s)
		return string(v)
	})
}
