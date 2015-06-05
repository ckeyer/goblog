package controllers

import (
	// "encoding/base64"
	"github.com/astaxie/beego"
	"github.com/ckeyer/goblog/models"
	// "log"
	"crypto/md5"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

type BlogController struct {
	beego.Controller
}

type BlogData struct {
	MsgCode    int      `form:"msgcode"`
	Title      string   `form:"titile"`
	Tags       []string `form:"tags[]"`
	BlogType   string   `form:"blog_type"`
	Password   string   `form:"password"`
	CommitSalt string   `form:"commit_value"`
	Summary    string   `form:"summary"`
	Content    string   `form:"content"`
}

func (this *BlogController) initData() {
	this.Data["STATIC_URL_JS"] = this.GetSession("STATIC_URL_JS")
	this.Data["STATIC_URL_IMG"] = this.GetSession("STATIC_URL_IMG")
	this.Data["STATIC_URL_CSS"] = this.GetSession("STATIC_URL_CSS")
}

// @router /v1/blog [get]
func (this *BlogController) ShowList() {
	this.initData()
	log.Println("// @router /v1/blog [get]")
	wp := models.NewWebPage("博客")
	wp.IncrViewCount()
	this.Ctx.WriteString("// @router /v1/blog [get]")
}

// @router /v1/blog::key([0-9]+).html [get]
func (this *BlogController) ShowBlog() {
	this.initData()
	log.Println("// @router /v1/blog::key([0-9]+).html [get]")
	s := this.Ctx.Input.Param(":key")
	this.Ctx.WriteString(s)
	this.Ctx.WriteString("// @router /v1/blog:key [get]")
}

func (this *BlogController) Post() {
	log.Println("")
}

// @router /v1/admin/blog [get]
func (this *BlogController) ShowEditList() {
	this.initData()
	log.Println("// @router /v1/admin/blog [get]")
	this.Ctx.WriteString("// @router /v1/admin/blog [get]")
}

// @router /v1/admin/blog:key([0-9]+).html [get]
func (this *BlogController) EditBlog() {
	this.initData()
	// s := this.Ctx.Input.Param(":key")
	// this.Ctx.WriteString(s)
	// this.Ctx.WriteString("// @router /v1/admin/blog [get]")
	this.TplNames = "blogEdit.tpl"
}

func (this *BlogController) AddBlog() {
	this.initData()
	// s := this.Ctx.Input.Param(":key")
	// this.Ctx.WriteString(s)
	// this.Ctx.WriteString("// @router /v1/admin/blog [get]")
	this.TplNames = "blogEdit.tpl"
}

// @router /v1/admin/blog/new [post]
func (this *BlogController) NewBlog() {

	blogpost := &BlogData{}
	err := this.ParseForm(blogpost)
	if err != nil {
		log.Println(err.Error())
		this.Ctx.WriteString(`{"msgcode":-1,"data":"` + err.Error() + `""}`)
	}
	log.Println("tags", blogpost.Tags)
	log.Println("post datat", blogpost)
	// log.Println(authCommit(blogpost.Password, blogpost.CommitSalt))
	switch blogpost.MsgCode {
	case 1:
		if !authCommit(blogpost.Password, blogpost.CommitSalt) {
			this.Ctx.WriteString(`{"msgcode":-2,"data":"auth error""}`)
			return
		}
		var tags []*models.Tag = make([]*models.Tag, len(blogpost.Tags))
		for i, v := range blogpost.Tags {
			tags[i] = models.GetTag(v)
		}
		blog := &models.Blog{
			Title:   blogpost.Title,
			Summary: blogpost.Summary,
			Content: blogpost.Content,
			Type:    blogpost.BlogType,
			Tags:    tags,
		}
		// log.Println("add blog", blog.Tags[0])
		if err := blog.Insert(); err != nil {
			this.Ctx.WriteString(`{"msgcode":-3,"data":"` + err.Error() + `""}`)
		} else {
			// log.Printf("add blog %v\n", blog.Tags[0])
			this.Ctx.WriteString(`{"msgcode":1}`)
			return
		}
	}
	this.Ctx.WriteString(`{"msgcode":0}`)
}

func authCommit(pwd, salt string) bool {
	fmtSalt := func(s string) string {
		var ns []int = make([]int, len(s))
		for i := 0; i < len(s); i++ {
			ns[i], _ = strconv.Atoi(string(s[i]))
		}
		nsort := sort.IntSlice(ns)
		nsort.Sort()
		a, b := 0, 0
		for i := 0; i < len(nsort)/2; i++ {
			a += nsort[i]
		}
		for i := len(nsort) / 2; i < len(nsort); i++ {
			b += nsort[i]
		}
		return fmt.Sprintf("%d%d", a, b)
	}
	// 获取字符串的MD5值
	getMD5 := func(s string) (md5s string) {
		f := md5.New()
		io.WriteString(f, s)
		md5s = fmt.Sprintf("%x", f.Sum(nil))
		return
	}

	pass := "wangcj" + fmtSalt(salt)
	return strings.ToLower(pwd) == strings.ToLower(getMD5(pass))
}
