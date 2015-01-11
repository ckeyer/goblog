package models

import (
	"github.com/astaxie/beego"
	"github.com/hoisie/redis"
	_ "log"
	"strconv"
	"strings"
)

type WebPage struct {
	Name string
	rc   redis.Client
}

func NewWebPage(name string) (w *WebPage) {
	w = &WebPage{}
	w.Name = name

	w.rc.Addr = beego.AppConfig.String("redis_addr")
	if "none" != name {
		w.rc.Sadd("WebPages", []byte(strings.ToUpper(w.Name)))
	}
	return
}
func (w *WebPage) GetPageTitle() string {
	return beego.AppConfig.String("host_name") + " - " + w.Name
}
func (w *WebPage) IncrViewCount() (count int) {
	key := strings.ToUpper(w.Name + "_ViewCount")
	v, _ := w.rc.Incr(key)
	count = int(v)
	return
}
func (w *WebPage) GetViewCount() (count int) {
	key := strings.ToUpper(w.Name + "_ViewCount")
	v, err := w.rc.Get(key)
	if err == nil {
		count, _ = strconv.Atoi(string(v))
	} else {
		w.rc.Set(key, []byte("100"))
		count = 100
	}
	return
}
func (w *WebPage) GetViewCountByName(name string) (count int) {
	key := strings.ToUpper(name + "_ViewCount")
	v, err := w.rc.Get(key)
	if err == nil {
		count, _ = strconv.Atoi(string(v))
	} else {
		w.rc.Set(key, []byte("100"))
		count = 100
	}
	return
}
func (w *WebPage) GetWebPages() map[string]int {
	key := "WebPages"
	s := make(map[string]int)
	vs, _ := w.rc.Smembers(key)
	for _, v := range vs {
		s["/"+strings.ToLower(string(v))] = w.GetViewCountByName(string(v))
	}
	return s
}
func (w *WebPage) GetWebPageCount() int {
	key := "WebPages"
	v, _ := w.rc.Scard(key)
	return v
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
