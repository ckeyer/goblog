/*
**/

package models

import (
	"log"
)

type Tag struct {
	Id       int
	ParentId int
	Name     string
	Blogs    []int
}

func NewTag(name string) *Tag {
	return &Tag{Name: name}
}

func GetTagByBlogId(blog_id int) *Tag {
	return &Tag{}
}

func (this *Tag) GetAllByName() {

}
