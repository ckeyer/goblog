/*
**/

package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Tag struct {
	Id       int64
	ParentId int    `orm:"default(0)"`
	Name     string `orm:"index;size(32);unique"`

	Blogs []*Blog `orm:"reverse(many)"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func NewTag(name string, parent int) *Tag {
	return &Tag{Name: name, ParentId: parent}
}
func GetTagById(id int64) (tag *Tag, err error) {
	o := orm.NewOrm()
	tag = &Tag{Id: id}

	err = o.Read(tag)

	return
}
func (this *Tag) Insert() bool {
	o := orm.NewOrm()

	id, err := o.Insert(this)
	if err == nil {
		this.Id = id
		return true
	}
	return false
}
func (this *Tag) Update(name string, parent int) bool {
	o := orm.NewOrm()
	this.Name = name
	this.ParentId = parent
	if _, err := o.Update(this); err == nil {
		return true
	}
	return false
}

func (this *Tag) Delete() bool {
	o := orm.NewOrm()
	if _, err := o.Delete(this); err == nil {
		return true
	}
	return false
}
func GetTagByBlogId(blog_id int) *Tag {
	return &Tag{}
}

func (this *Tag) GetAllByName() {

}
