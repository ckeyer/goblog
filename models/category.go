package models

import (
	. "github.com/ckeyer/goblog/modules"
)

func GetAllCategory() (ts []string) {

	cats := MyPool.GetCategories()
	ts = make([]string, cats.Len())
	for i, v := range cats.GetAll() {
		ts[i] = v.Name
	}

	return
}

func GetBlogsByCategory(name string) (bs Blogs) {
	tags := MyPool.GetCategories()
	return tags.GetByName(name).Blogs
}
