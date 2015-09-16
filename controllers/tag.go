package controllers

import (
	"github.com/ckeyer/goblog/models"
)

type TagController struct {
	BaseController
}

func (c *TagController) Get() {
	name := c.GetString("t")

	if tg := models.GetBlogsByTag(name); tg != nil {
		c.Data["Blogs"] = tg
	} else {
		log.Errorf("Tag Errer %s", name)
		c.Data["Blogs"] = models.GetBlogs(0, 10)
	}

	c.Data["LastestBlogs"] = models.GetBlogs(0, 5)
	c.Data["Tags"] = models.GetAllTags()
	c.Data["Category"] = models.GetAllCategory()
	c.Data["MonthBlog"] = models.GetAllMonth()

	c.LayoutSections["Sidebar"] = "sidebar.tpl"

	c.TplNames = "list.tpl"
}
