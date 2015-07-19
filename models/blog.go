package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Blog struct {
	Id      int64
	Title   string `orm:"size(256)"`
	Page    int    `orm:"default(0)"`
	Summary string `orm:"type(text); null"`
	Content string `orm:"type(text); null"`
	Type    string `orm:"size(32)"`
	Status  int    `orm:"default(0)"`

	Tags []*Tag `orm:"-"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}
type BlogsMonth struct {
	Month     string
	BlogCount int
}

func (this *Blog) AddTagName(tag_name string) {
	tag := &Tag{Name: tag_name}
	this.Tags = append(this.Tags, tag)
}
func GetBlogById(id int64) (b *Blog, err error) {
	o := orm.NewOrm()
	b = &Blog{Id: id}

	err = o.Read(b)
	if err != nil {
		log.Error(err.Error())
	}
	b.getTags()
	log.Debug("%V\n", b)
	return
}
func (this *Blog) Insert() error {
	o := orm.NewOrm()

	id, err := o.Insert(this)
	if err == nil {
		this.Id = id
	}
	this.getTags()
	return err
}
func (this *Blog) getTags() {
	o := orm.NewOrm()
	sql := "select tag.* from tag, blog_tag_relation as bt  where tag.id = bt.tag_id and bt.blog_id = ?"
	o.Raw(sql, this.Id).QueryRows(&this.Tags)
}
func (this *Blog) WriteToDB() (e error) {
	o := orm.NewOrm()
	e = o.Begin()
	// BEGIN
	e = this.Insert()
	if e == nil {
		for _, v := range this.Tags {
			log.Debug("%#v", v)
			e = v.Get()
			if e != nil {
				log.Error(e.Error())
				break
			}
			InsertBlogTagRelation(this, v)
		}
	}
	// END
	if e != nil {
		e = o.Rollback()
	} else {
		e = o.Commit()
	}
	return
}
func (this *Blog) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(this)
	return err
}
func (this *Blog) Delete() error {
	o := orm.NewOrm()
	_, err := delBlogInRelation(this)
	if err != nil {
		return err
	}
	_, err = o.Delete(this)
	return err
}

//用于格式化时间字段
func getTimeString(ttt time.Time) string {
	return (fmt.Sprintf("%d-%02d-%02d %02d:%02d",
		ttt.Year(), ttt.Month(), ttt.Day(),
		ttt.Hour(), ttt.Minute()))
}

// func (this time.Time) MarshalText() ([]byte, error) {
// 	return []byte(fmt.Sprintf("%d-%02d-%02d %02d:%02d",
// 		this.Created.Year(), this.Created.Month(), this.Created.Day(),
// 		this.Created.Hour(), this.Created.Minute())), nil
// }
func (this *Blog) ToJSON() (s string) {
	b, e := json.Marshal(this)
	if e != nil {
		return ""
	}
	fmt.Println(string(b))
	var mp map[string]interface{}
	e = json.Unmarshal(b, &mp)
	if e != nil {
		return ""
	}
	mp["Created"] = getTimeString(this.Created)
	mp["Updated"] = getTimeString(this.Updated)
	b, e = json.Marshal(&mp)
	if e != nil {
		return ""
	}
	return string(b)
}

func (this *Blog) Read() {
}
func (this *Blog) ReadBlogByID(id_ string) error {
	return nil
}

func (this *Blog) GetNextBlog() (b *Blog, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(b)
	qs.Filter("blog__id__lt", this.Id)
	qs.Limit(1)
	qs.OrderBy("-blog__id")
	// ORDER BY id ASC, profile.age DESC
	err = qs.One(b)
	return
}
func (this *Blog) GetPreviousBlog() (b *Blog, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(b)
	qs.Filter("blog__id__gt", this.Id)
	qs.Limit(1)
	err = qs.One(b)
	return
}
func GetBlogs(start, count int) (bs []*Blog) {
	o := orm.NewOrm()
	// res := make(orm.Params)
	sql := "select * from blog order by created desc limit ?,?"
	num, err := o.Raw(sql, start, start+count).QueryRows(&bs)
	if num == 0 || err != nil {
		log.Error("Error Getblogs :Get :%d,Error: %v\n", num, err)
	}
	for _, v := range bs {
		v.getTags()
	}
	return
}

// 获取按月的文章数
func GetBlogsMonth(cols int) (bs []*BlogsMonth) {
	o := orm.NewOrm()
	sql := "select DATE_FORMAT(created,'%Y-%m') as month,count(id) as blog_count from blog   group by month   order by month limit 0,?"

	num, err := o.Raw(sql, cols).QueryRows(&bs)
	if num == 0 || err != nil {
		log.Error("Error Getblogs :Get :%d,Error: %v\n", num, err)
	}
	return
}
