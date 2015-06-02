package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

const (
	RelationFavourite = "favourite"
)

type BlogTagRelation struct {
	Id         int64
	BlogId     int64     `orm:"index"`
	TagId      int64     `orm:"index"`
	ActionType string    `orm:"default('-')"`
	Created    time.Time `orm:"auto_now_add;type(datetime)"`
	Updated    time.Time `orm:"auto_now;type(datetime)"`
}

func InsertBlogTags(blogid, tagid int64) *BlogTagRelation {
	blogtag := &BlogTagRelation{
		BlogId: blogid,
		TagId:  tagid,
	}
	o := orm.NewOrm()
	id, err := o.Insert(blogtag)

	if err != nil {
		o.Read(blogtag)
	} else {
		blogtag.Id = id
	}
	return blogtag
}
func DelBlogTags(blogid int64, tags []int64) int {
	o := orm.NewOrm()
	count := 0
	for _, v := range tags {
		bt := &BlogTagRelation{
			BlogId: blogid,
			TagId:  v,
		}
		o.Delete(bt)
	}
	return count
}
