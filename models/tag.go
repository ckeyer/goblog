/*
**/

package models

import (
	"time"
)

type Tag struct {
	Id       int
	ParentId int     `orm:"default(1)"`
	Name     string  `orm:"index;size(32)"`
	Blogs    []*Blog `orm:"reverse(many)"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func NewTag(name string) *Tag {
	return &Tag{Name: name}
}

func GetTagByBlogId(blog_id int) *Tag {
	return &Tag{}
}

func (this *Tag) GetAllByName() {

}
