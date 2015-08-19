package main

import (
	"encoding/base64"

	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/session"
	"fmt"
	"strings"
	"time"

	"github.com/ckeyer/goblog/models"
	"github.com/ckeyer/goblog/routers"
)

func main() {
	db_init()
	server_init()
	config_init()
	routers.Error_init()
	beego.Run()
}

func server_init() {
	// beego.EnableHttpTLS = true
	// beego.HttpsPort = 443
	beego.SessionOn = true
	beego.SessionGCMaxLifetime = 2
	beego.SessionName = "CkeyerSession"
	beego.BeegoServerName = "ckeyer Server 1.0"
	// beego.HttpCertFile = "/var/www/https/cert.pem"
	// beego.HttpKeyFile = "/var/www/https/key.pem"
}
func db_init() {
	models.RegistDB()
}
func config_init() {
	beego.TemplateLeft = "<%"
	beego.TemplateRight = "%>"

	beego.AddFuncMap("FMT_DATETIME", func(t time.Time) string {
		return fmt.Sprintf("%v-%02d-%02d  %02d:%02d",
			t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
	})
	beego.AddFuncMap("FMT_DATE_MONTH", func(t time.Time) string {
		return fmt.Sprintf("%v-%02d", t.Year(), t.Month())
	})
	beego.AddFuncMap("STATIC_URL", func() string {
		return beego.AppConfig.String("static_rul")
	})
	beego.AddFuncMap("SITE_URL", func() string {
		return beego.AppConfig.String("site_url")
	})
	beego.AddFuncMap("STATIC_URL_JS", func() string {
		return beego.AppConfig.String("static_url_js_ssl")
	})
	beego.AddFuncMap("STATIC_URL_CSS", func() string {
		return beego.AppConfig.String("static_url_css_ssl")
	})
	beego.AddFuncMap("STATIC_URL_IMG", func() string {
		return beego.AppConfig.String("static_url_img_ssl")
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
