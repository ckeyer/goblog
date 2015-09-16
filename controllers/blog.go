package controllers

import (
	"github.com/ckeyer/goblog/models"
)

type BlogController struct {
	BaseController
}

// (b *BaseController)GetBlogs ...
func (b *BlogController) Get() {
	b.Data["LastestBlogs"] = models.GetBlogs(0, 5)
	b.Data["Tags"] = models.GetAllTags()
	b.Data["Category"] = models.GetAllCategory()
	b.Data["MonthBlog"] = models.GetAllMonth()

	name := b.Ctx.Input.Param(":name")
	blog := models.GetBlog(name)
	if blog == nil {
		log.Debug("name: ", name)
		return
	}
	b.Data["Blog"] = blog
	b.Data["BContent"] = string(blog.Content)

	b.LayoutSections["Sidebar"] = "sidebar.tpl"
	b.LayoutSections["Duoshuo"] = "duoshuo.tpl"
	b.TplNames = "show.tpl"
}
