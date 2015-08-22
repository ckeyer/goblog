package controllers

import (
	"crypto/md5"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/ckeyer/goblog/models"
)

type BlogController struct {
	BaseController
}

type BlogData struct {
	Code       int      `form:"code"`
	Title      string   `form:"titile"`
	Tags       []string `form:"tags[]"`
	BlogType   string   `form:"blog_type"`
	Password   string   `form:"password"`
	CommitSalt string   `form:"commit_value"`
	Summary    string   `form:"summary"`
	Content    string   `form:"content"`
}
type BlogUpJson struct {
}

// @router /v1/blog [get]
func (this *BlogController) ShowList() {
	log.Debug("// @router /v1/blog [get]")
	wp := models.NewWebPage("博客")
	wp.IncrViewCount()
	this.Ctx.WriteString("// @router /v1/blog [get]")
}

// @router /v1/blog::key([0-9]+).html [get]
func (this *BlogController) ShowBlog() {

	skey := this.Ctx.Input.Param(":key")
	key, err := strconv.Atoi(skey)
	if err != nil {
		this.Ctx.WriteString("Error")
		return
	}
	blog, err := models.GetBlogById(int64(key))
	if err != nil {
		this.Ctx.WriteString("Error")
		return
	}
	this.Data["Article"] = blog
	this.TplNames = "blog.tpl"
}

func (this *BlogController) Post() {
	log.Debug("Post oKKKKKKKK")
	code, _ := this.GetInt32("code")
	switch code {
	case 1:
		id, err := this.GetInt64("id")
		if err != nil {
			this.Ctx.WriteString(`{"code":-2,"data":"Get err id "}`)
			break
		}
		blog, err := models.GetBlogById(id)
		if err != nil {
			this.Ctx.WriteString(`{"code":-3,"data":"Get err blog "}`)
			break
		}
		log.Debug("...")
		this.Ctx.WriteString(`{"code":-3,"data":` + blog.ToJSON() + `}`)
		return
	default:
		this.Ctx.WriteString(`{"code":-1,"data":"Crying"}`)
	}
}

// @router /v1/admin/blog [get]
func (this *BlogController) ShowEditList() {
	log.Debug("// @router /v1/admin/blog [get]")
	this.Ctx.WriteString("// @router /v1/admin/blog [get]")
}

// @router /v1/admin/blog:key([0-9]+).html [get]
func (this *BlogController) EditBlog() {
	// s := this.Ctx.Input.Param(":key")
	// this.Ctx.WriteString(s)
	// this.Ctx.WriteString("// @router /v1/admin/blog [get]")
	this.TplNames = "blogEdit.tpl"
}

func (this *BlogController) AddBlog() {
	// s := this.Ctx.Input.Param(":key")
	// this.Ctx.WriteString(s)
	// this.Ctx.WriteString("// @router /v1/admin/blog [get]")
	this.Data["CusStyles"] = `<link rel="Stylesheet" type="text/css" href="` + website.CssUrl + `jquery-ui.1.11.3.min.css" />
<link rel="Stylesheet" type="text/css" href="` + website.CssUrl + `jHtmlArea.css" />`

	this.Data["CusScripts"] = `<script type="text/javascript" src="` + website.JsUrl + `jquery-ui.1.11.3.min.js"></script>
<script type="text/javascript" src="` + website.JsUrl + `jHtmlArea-0.8.min.js"></script>
<script type="text/javascript" src="` + website.JsUrl + `edit_blog.js"></script>`

	this.TplNames = "blogEdit.tpl"
}

// @router /v1/admin/blog/new [post]
func (this *BlogController) NewBlog() {

	blogpost := &BlogData{}
	err := this.ParseForm(blogpost)
	if err != nil {
		log.Debug(err.Error())
		this.Ctx.WriteString(`{"code":-1,"data":"` + err.Error() + `""}`)
	}
	log.Debug("tags", blogpost.Tags)
	log.Debug("post datat", blogpost)
	// log.Debug(authCommit(blogpost.Password, blogpost.CommitSalt))
	switch blogpost.Code {
	case 1:
		if !authCommit(blogpost.Password, blogpost.CommitSalt) {
			this.Ctx.WriteString(`{"code":-2,"data":"auth error""}`)
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
		// log.Debug("add blog", blog.Tags[0])
		if err := blog.WriteToDB(); err != nil {
			this.Ctx.WriteString(`{"code":-3,"data":"` + err.Error() + `""}`)
		} else {
			// log.Printf("add blog %v\n", blog.Tags[0])
			this.Ctx.WriteString(`{"code":1}`)
			return
		}
	}
	this.Ctx.WriteString(`{"code":0}`)
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
