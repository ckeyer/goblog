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
	ActionType string    `orm:"size(2);default(-)"`
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

func delBlog(blogid int64) {
	o := orm.NewOrm()

	qs := o.QueryTable("blog_tag_relation")

	// 也可以直接使用对象作为表名
	btr := new(BlogTagRelation)
	qs = o.QueryTable(btr)

	qs.Filter("blog_id", blogid)
	btrs, err := qs.All()
	if err != nil {
		log.Println("error 1", err.Error())
		return
	}

	for i, v := range btrs {
		if num, err := o.Delete(&BlogTagRelation{BlogId: v}); err == nil {
			fmt.Println(num)
		}
	}

}

// if num, err := o.Delete(&User{Id: 1}); err == nil {
//     fmt.Println(num)
// }
