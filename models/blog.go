package models

import (
	"container/list"
	"github.com/astaxie/beego/orm"
	"time"
)

type Blog struct {
	Id      int64
	Title   string `orm:"size(32)"`
	Page    int    `orm:"default(0)"`
	Summary string
	Content string
	Type    string `orm:"size(12)"`
	Status  int    `orm:"default(0)"`

	Tags []*Tag `orm:"reverse(many)"`

	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func GetBlogById(id int64) (b *Blog, err error) {
	o := orm.NewOrm()
	b = &Blog{Id: id}

	err = o.Read(b)
	return
}
func NewBlog() *Blog {
	return &Blog{}
}
func NewBlogByValue(title string, content string, summary string) (b *Blog) {
	b = &Blog{
		Title:   title,
		Summary: summary,
		Content: content,
	}
	return b
}
func (this *Blog) Insert() error {
	o := orm.NewOrm()

	id, err := o.Insert(this)
	if err == nil {
		this.Id = id
	}
	return err
}

func (this *Blog) Update() error {
	o := orm.NewOrm()
	_, err := o.Update(this)
	return err
}
func (this *Blog) Delete() error {
	o := orm.NewOrm()
	_, err := o.Delete(this)
	return err
}
func (this *Blog) AddTag(t *Tag) {
	this.Tags = append(this.Tags, t)
}

func (this *Blog) ToMap() map[string]string {
	bm := make(map[string]string)
	// 	bm["id"] = strconv.Itoa(this.Id)
	// 	bm["title"] = this.Title
	// 	bm["page"] = strconv.Itoa(this.Page)
	// 	bm["summary"] = this.Summary
	// 	bm["author_id"] = strconv.Itoa(this.AuthorID)
	// 	bm["content"] = this.Content
	// 	bm["status"] = strconv.Itoa(this.Status)
	// 	bm["created"] = this.CreatedTime.String()
	// 	bm["updated"] = this.UpdateTme.String()
	return bm
}
func (this *Blog) ToJSON() (s string) {
	return ""
}

func (this *Blog) Read() {
}
func (this *Blog) ReadBlogByID(id_ string) error {
	return nil
}

// func (this *Blog) readTags() error {
// 	this.connectDB()
// 	defer this.close()
// 	sqlStr := "select tb_tag.* from tb_art_tag,tb_tag where tb_tag.id=tb_art_tag.tag_id "

// 	defer func() {
// 		if e := recover(); e != nil {
// 			log.Println("Error:", e)
// 		}
// 	}()
// 	rows, err := this.db.Query(sqlStr+" and tb_art_tag.art_id=? ", strconv.Itoa(this.ID))
// 	if err != nil {
// 		return err
// 	}
// 	var id int
// 	var parent_id int
// 	var name string

// 	for rows.Next() {
// 		if err := rows.Scan(&id, &parent_id, &name); err == nil {
// 			this.Tags = append(this.Tags, NewTag(id, name, parent_id, 0))
// 		} else {
// 			return err
// 		}
// 	}

// 	return nil
// }
// func (this *Blog) GetHotTags() (ts []*Tag) {
// 	this.connectDB()
// 	defer this.close()
// 	sqlStr := "select tb_tag.*, count(tb1.art_id) as count_art_id " +
// 		"from tb_art_tag as tb1 , tb_tag where tb_tag.id = tb1.tag_id " +
// 		"group by tb1.tag_id order by count_art_id desc limit 0,7"

// 	rows, err := this.db.Query(sqlStr)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return
// 	}
// 	var id int
// 	var parent_id int
// 	var name string
// 	var count_art_id int

// 	for rows.Next() {
// 		if err := rows.Scan(&id, &parent_id, &name, &count_art_id); err == nil {
// 			ts = append(ts, NewTag(id, name, parent_id, count_art_id))
// 		} else {
// 			log.Println(err.Error())
// 		}
// 	}
// 	return
// }
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

func (this *Blog) GetBlogList(start, stop int) (bs *list.List) {
	return
}
func (this *Blog) GetBlogs(start, stop int) (bs []*Blog) {
	return
}
func (this *Blog) GetBlogsByTagId(tag_id, start, stop int) (bs []*Blog) {
	return
}
