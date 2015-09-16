package controllers

import (
	"github.com/ckeyer/goblog/models"
)

type CategoryController struct {
	BaseController
}

func (c *CategoryController) Get() {
	name := c.GetString("c")

	if tg := models.GetBlogsByCategory(name); tg != nil {
		c.Data["Blogs"] = tg
	} else {
		log.Errorf("Category Errer %s", name)
		c.Data["Blogs"] = models.GetBlogs(0, 10)
	}

	c.Data["LastestBlogs"] = models.GetBlogs(0, 5)
	c.Data["Tags"] = models.GetAllTags()
	c.Data["Category"] = models.GetAllCategory()
	c.Data["MonthBlog"] = models.GetAllMonth()

	c.LayoutSections["Sidebar"] = "sidebar.tpl"

	c.TplNames = "list.tpl"
}
