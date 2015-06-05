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
	Blog       *Blog     `orm:"rel(fk)"`
	Tag        *Tag      `orm:"rel(fk)"`
	ActionType string    `orm:"size(2);default(-)"`
	Created    time.Time `orm:"auto_now_add;type(datetime)"`
	Updated    time.Time `orm:"auto_now;type(datetime)"`
}

func InsertBlogTags(b *Blog, t *Tag) *BlogTagRelation {
	blogtag := &BlogTagRelation{
		Blog: b,
		Tag:  t,
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
func DelBlogTags(b *Blog, tags []*Tag) int {
	o := orm.NewOrm()
	count := 0
	for _, v := range tags {
		bt := &BlogTagRelation{
			Blog: b,
			Tag:  v,
		}
		o.Delete(bt)
	}
	return count
}

func delBlogInRelation(b *Blog) (int64, error) {
	o := orm.NewOrm()

	btr := new(BlogTagRelation)
	qs := o.QueryTable(btr)

	qs.Filter("blog_id", b.Id)
	count, err := qs.Delete()
	return count, err
}

// if num, err := o.Delete(&User{Id: 1}); err == nil {
//     fmt.Println(num)
// }
