package modules

import (
	"testing"
)

// TestGetFiles ...
func TestGetFiles(t *testing.T)  {
	err := LoadBlogs("../blog/*.md")
	if err != nil {
		t.Error(err)
	}
	tags:=pool.GetTags()
	l:=tags.Len()
	log.Debug("tags len =",l)
	for i,v:= range tags.GetAll() {
		log.Debugf("%d %#v",i ,*v)
	}

	cat := pool.GetCategory()
	l = cat.Len()
	log.Debug("category len =",l)
	for i,v:= range cat .GetAll() {
		log.Debugf("%d %#v",i ,*v)
	}

	mblog := pool.GetMonthBlog()
	l = mblog.Len()
	log.Debug("monthblogs len =",l)
	for i,v:= range mblog.GetAll() {
		log.Debugf("%d %#v",i ,*v)
	}
}
