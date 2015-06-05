/*
**/

package models

import (
	// "strings"
	"github.com/astaxie/beego/orm"
	"time"
)

type Tag struct {
	Id       int64
	ParentId int    `orm:"default(0)"`
	Name     string `orm:"pk;size(32);unique"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func NewTag(name string, parent int) *Tag {
	return &Tag{Name: name, ParentId: parent}
}
func GetTag(tag_name string) *Tag {
	o := orm.NewOrm()
	tag := &Tag{Name: tag_name}
	_, id, err := o.ReadOrCreate(tag, tag_name)
	o.ReadOrCreate(tag, "Name")
	if err != nil {
		return nil
	}
	tag.Id = id
	return tag
}
func (this *Tag) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(this)
	if err != nil {
		log.Println(err.Error())
	}
	return err
}
func (this *Tag) Delete() error {
	o := orm.NewOrm()
	_, err := o.Delete(this)
	return err
}
