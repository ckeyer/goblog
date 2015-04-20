package models

import (
	"container/list"
	// "github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// "strconv"
	"time"
)

type Blog struct {
	Id       int
	Title    string
	Page     int
	AuthorId int
	Summary  string
	Content  string
	Status   int
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`

	Tags []*Tag
}

func NewBlog() (b *Blog) {
	b = &Blog{}
	return b
}
func NewBlogByValue(id int, title string, page int, author_id int, content string, summary string, status int) (b *Blog) {
	b = &Blog{
		Id:       id,
		Title:    title,
		Page:     page,
		AuthorId: author_id,
		Summary:  summary,
		Content:  content,
		Status:   status,
	}
	return b
}
func (this *Blog) insertToDB() {

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
func (this *Blog) GetNextBlog() (b *Blog) {
	return
}
func (this *Blog) GetPreviousBlog() (b *Blog) {
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
