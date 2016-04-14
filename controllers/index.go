package controllers

import (
	"github.com/ckeyer/goblog/models"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {
	page, _ := c.GetInt("page")
	step := 15

	c.Data["Blogs"] = models.GetBlogs(page*step, (page+1)*step)
	c.Data["LastestBlogs"] = models.GetBlogs(0, 5)
	c.Data["Tags"] = models.GetAllTags()
	c.Data["Category"] = models.GetAllCategory()
	c.Data["MonthBlog"] = models.GetAllMonth()

	c.LayoutSections["Sidebar"] = "sidebar.tpl"

	c.TplName = "list.tpl"
}
