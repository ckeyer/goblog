package models

import (
	_ "log"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type WebPage struct {
	Name string
}

func NewWebPage(name string) (w *WebPage) {
	w = &WebPage{}
	w.Name = name

	if "none" != name {
		rc.SAdd("WebPages", (strings.ToUpper(w.Name)))
	}
	return
}
func (w *WebPage) GetPageTitle() string {
	return beego.AppConfig.String("host_name") + " - " + w.Name
}
func (w *WebPage) IncrViewCount() (count int) {
	key := strings.ToUpper(w.Name + "_ViewCount")
	v := rc.Incr(key)
	count = int(v.Val())
	return
}
func (w *WebPage) GetViewCount() (count int) {
	key := strings.ToUpper(w.Name + "_ViewCount")
	v := rc.Get(key)
	if v.Err() == nil {
		count, _ = strconv.Atoi(v.Val())
	} else {
		rc.Set(key, "100", 0)
		count = 100
	}
	return
}
func (w *WebPage) GetViewCountByName(name string) (count int) {
	key := strings.ToUpper(name + "_ViewCount")
	v := rc.Get(key)
	if v.Err() == nil {
		count, _ = strconv.Atoi((v.Val()))
	} else {
		rc.Set(key, ("100"), 0)
		count = 100
	}
	return
}
func (w *WebPage) GetWebPages() map[string]int {
	key := "WebPages"
	s := make(map[string]int)
	vs := rc.SMembers(key)
	for _, v := range vs.Val() {
		s["/"+strings.ToLower(string(v))] = w.GetViewCountByName(string(v))
	}
	return s
}
func (w *WebPage) GetWebPageCount() int {
	key := "WebPages"
	v := rc.SCard(key)
	return int(v.Val())
}
func (w *WebPage) GetImgHost() (s string) {
	s = beego.AppConfig.String("img_host")
	return s
}
func (w *WebPage) GetStaticHost() (s string) {
	s = beego.AppConfig.String("static_host")
	return
}

func (w *WebPage) Test() (count int) {
	count = 0
	return
}
