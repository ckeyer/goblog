package models

import (
	. "github.com/ckeyer/goblog/modules"
)

// GetAllTags ...
func GetAllTags() (ts []string) {

	tags := MyPool.GetTags()
	ts = make([]string, tags.Len())
	for i, v := range tags.GetAll() {
		ts[i] = v.Name
	}

	return
}

// GetBlogsByTag ...
func GetBlogsByTag(name string) (bs Blogs) {
	tags := MyPool.GetTags()
	return tags.GetByName(name).Blogs
}
