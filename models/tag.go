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
	Name     string `orm:"size(32);unique"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

// func NewTag(name string, parent int) *Tag {
// 	return &Tag{Name: name, ParentId: parent}
// }

func GetTag(tag_name string) *Tag {
	tag := &Tag{Name: tag_name}
	tag.getByName()
	return tag
}
func (this *Tag) Get() error {
	if this.Name != "" {
		return this.getByName()
	} else {
		return this.getById()
	}
}
func (this *Tag) getById() error {
	o := orm.NewOrm()
	return o.Read(this)
}
func (this *Tag) getByName() error {
	o := orm.NewOrm()
	_, id, err := o.ReadOrCreate(this, "name")
	log.Println("#$@#################@", id, err)
	o.ReadOrCreate(this, "Name")
	if err == nil {
		this.Id = id
	}
	return err

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
func GetHotTags(max int) (tags []*Tag) {
	o := orm.NewOrm()
	// num, err :=
	o.Raw("select * from blog_tag_relation as bt group by bt.tag_id  order by count(bt.blog_id) desc", 1).QueryRows(tags)
	return
}
