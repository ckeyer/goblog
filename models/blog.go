package models

import (
	"container/list"
	"database/sql"
	// "errors"
	// "fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

type Tag struct {
	ID       int
	Name     string
	ParentID int

	ArtCount int
}
type Blog struct {
	ID          int
	Title       string
	Page        int
	AuthorID    int
	Summary     string
	Content     string
	Status      int
	CreatedTime string
	UpdateTme   string

	Tags []*Tag

	db *sql.DB
}

func NewBlog() (b *Blog) {
	b = &Blog{}
	return b
}
func NewTag(id int, name string, parent_id int, art_count int) (tag *Tag) {
	tag = &Tag{
		ID:       id,
		Name:     name,
		ParentID: parent_id,
		ArtCount: art_count,
	}
	return
}
func NewBlogByValue(id int, title string, page int, author_id int, content string, summary string, status int, created string, updated string) (b *Blog) {
	b = &Blog{
		ID:          id,
		Title:       title,
		Page:        page,
		AuthorID:    author_id,
		Summary:     summary,
		Content:     content,
		Status:      status,
		CreatedTime: created,
		UpdateTme:   updated,
	}
	return b
}
func (this *Blog) connectDB() {
	if nil == this.db || this.db.Ping() != nil {
		dbStr := beego.AppConfig.String("sql_db")
		connStr := beego.AppConfig.String("sql_conn_str")
		var err error
		this.db, err = sql.Open(dbStr, connStr)

		if err != nil {
			log.Println("database initialize error : ", err.Error())
			// this.db = db
		}
	}
}
func (this *Blog) close() {
	if nil != this.db {
		this.db.Close()
	}
}

func (this *Blog) ToMap() map[string]string {
	bm := make(map[string]string)
	bm["id"] = strconv.Itoa(this.ID)
	bm["title"] = this.Title
	bm["page"] = strconv.Itoa(this.Page)
	bm["summary"] = this.Summary
	bm["author_id"] = strconv.Itoa(this.AuthorID)
	bm["content"] = this.Content
	bm["status"] = strconv.Itoa(this.Status)
	bm["created"] = this.CreatedTime
	bm["updated"] = this.UpdateTme
	return bm
}
func (this *Blog) ToJSON() (s string) {
	s = `{"id":` + strconv.Itoa(this.ID) + `,` +
		`"title":"` + this.Title + `",` +
		`"page":` + strconv.Itoa(this.Page) + `,` +
		`"summary":"` + this.Summary + `",` +
		`"author_id":` + strconv.Itoa(this.AuthorID) + `,` +
		`"content":"` + this.Content + `",` +
		`"status":` + strconv.Itoa(this.Status) + `,` +
		`"created":"` + this.CreatedTime + `",` +
		`"updated":"` + this.UpdateTme + `"}`
	return
}
func (this *Blog) Read() {
	this.connectDB()
	defer this.close()
	sqlStr := "select id, title, page, author_id, summary, content, status, DATE_FORMAT(created,'%Y年%c月%d日 %T'), DATE_FORMAT(updated,'%Y年%c月%d日 %T') from tb_article "

	defer func() {
		if e := recover(); e != nil {
			log.Println("Error:", e)
		}
	}()

	rows, err := this.db.Query(sqlStr)
	if err != nil {
		log.Println("database select error: ", err.Error())
		return
	}

	var id int
	var title string
	var page int
	var author_id int
	var summary string
	var content string
	var status int
	var created string
	var updated string

	for rows.Next() {
		if err := rows.Scan(&id, &title, &page, &author_id, &summary, &content, &status, &created, &updated); err == nil {
			log.Println(id, title, page, author_id, summary, content, status, created, updated)
		} else {
			log.Println(err.Error())
		}
	}
}
func (this *Blog) ReadBlogByID(id_ string) error {
	this.connectDB()
	defer this.close()
	sqlStr := "select id, title, page, author_id, summary, content, status, DATE_FORMAT(created,'%Y年%c月%d日 %T'), DATE_FORMAT(updated,'%Y年%c月%d日 %T') from tb_article "

	defer func() {
		if e := recover(); e != nil {
			log.Println("Error:", e)
			// err := errors.New(fmt.Sprint(e))
			// return ?err
		}
	}()

	rows := this.db.QueryRow(sqlStr+" where id=?", id_)

	var id int
	var title string
	var page int
	var author_id int
	var summary string
	var content string
	var status int
	var created string
	var updated string

	if err := rows.Scan(&id, &title, &page, &author_id, &summary, &content, &status, &created, &updated); err == nil {
		// log.Println(id, title, page, author_id, summary, content, status, created, updated)
		this.ID = id
		this.Title = title
		this.Page = page
		this.AuthorID = author_id
		this.Summary = summary
		this.Content = content
		this.Status = status
		this.CreatedTime = created
		this.UpdateTme = updated
	} else {
		return err
	}
	if err := this.readTags(); err != nil {
		return err
	}
	return nil
}
func (this *Blog) readTags() error {
	this.connectDB()
	defer this.close()
	sqlStr := "select tb_tag.* from tb_art_tag,tb_tag where tb_tag.id=tb_art_tag.tag_id "

	defer func() {
		if e := recover(); e != nil {
			log.Println("Error:", e)
		}
	}()
	rows, err := this.db.Query(sqlStr+" and tb_art_tag.art_id=? ", strconv.Itoa(this.ID))
	if err != nil {
		return err
	}
	var id int
	var parent_id int
	var name string

	for rows.Next() {
		if err := rows.Scan(&id, &parent_id, &name); err == nil {
			this.Tags = append(this.Tags, NewTag(id, name, parent_id, 0))
		} else {
			return err
		}
	}

	return nil
}
func (this *Blog) GetHotTags() (ts []*Tag) {
	this.connectDB()
	defer this.close()
	sqlStr := "select tb_tag.*, count(tb1.art_id) as count_art_id " +
		"from tb_art_tag as tb1 , tb_tag where tb_tag.id = tb1.tag_id " +
		"group by tb1.tag_id order by count_art_id desc limit 0,7"

	rows, err := this.db.Query(sqlStr)
	if err != nil {
		log.Println(err.Error())
		return
	}
	var id int
	var parent_id int
	var name string
	var count_art_id int

	for rows.Next() {
		if err := rows.Scan(&id, &parent_id, &name, &count_art_id); err == nil {
			ts = append(ts, NewTag(id, name, parent_id, count_art_id))
		} else {
			log.Println(err.Error())
		}
	}
	return
}
func (this *Blog) GetNextBlog() (b *Blog) {
	this.connectDB()
	defer this.close()
	sqlStr := "select id,title from tb_article where id<? order by id desc limit 0,1"

	defer func() {
		if e := recover(); e != nil {
			log.Println("Error:", e)
			// err := errors.New(fmt.Sprint(e))
			// return ?err
		}
	}()

	rows := this.db.QueryRow(sqlStr, strconv.Itoa(this.ID))

	var id int
	var title string

	if err := rows.Scan(&id, &title); err == nil {
		// log.Println(id, title, page, author_id, summary, content, status, created, updated)
		b = NewBlog()
		b.ID = id
		b.Title = title
	} else {
		return nil
	}
	return
}
func (this *Blog) GetPreviousBlog() (b *Blog) {
	this.connectDB()
	defer this.close()
	sqlStr := "select id,title from tb_article where id>? limit 0,1"

	defer func() {
		if e := recover(); e != nil {
			log.Println("Error:", e)
		}
	}()

	rows := this.db.QueryRow(sqlStr, strconv.Itoa(this.ID))

	var id int
	var title string

	if err := rows.Scan(&id, &title); err == nil {
		// log.Println(id, title, page, author_id, summary, content, status, created, updated)
		b = NewBlog()
		b.ID = id
		b.Title = title
	} else {
		return nil
	}
	return
}

func (this *Blog) GetBlogList(start, stop int) (bs *list.List) {
	this.connectDB()
	defer this.close()
	sqlStr := "select id, title, page, author_id, content, summary, status, " +
		"DATE_FORMAT(created,'%Y年%c月%d日 %T')," +
		" DATE_FORMAT(updated,'%Y年%c月%d日 %T') " +
		"from tb_article order by id desc limit " +
		strconv.Itoa(start) + "," + strconv.Itoa(stop)
	defer func() {
		if e := recover(); e != nil {
			log.Println("Error:", e)
		}
	}()
	rows, err := this.db.Query(sqlStr)
	if err != nil {
		log.Println("database select error: ", err.Error())
		return nil
	}
	defer rows.Close()

	var id int
	var title string
	var page int
	var author_id int
	var summary string
	var content string
	var status int
	var created string
	var updated string

	bs = list.New()
	for rows.Next() {
		if err := rows.Scan(&id, &title, &page, &author_id, &summary, &content, &status, &created, &updated); err == nil {
			// log.Println(id, title, page, author_id, summary, content, status, created, updated)
			bs.PushBack(NewBlogByValue(id, title, page, author_id, summary, content, status, created, updated))
		} else {
			log.Println(err.Error())
		}
	}
	return bs
}
func (this *Blog) GetBlogs(start, stop int) (bs []*Blog) {
	this.connectDB()
	defer this.close()
	sqlStr := "select id, title, page, author_id, content, summary, status, " +
		"DATE_FORMAT(created,'%Y年%c月%d日 %T')," +
		" DATE_FORMAT(updated,'%Y年%c月%d日 %T') " +
		"from tb_article order by id desc limit " +
		strconv.Itoa(start) + "," + strconv.Itoa(stop)
	defer func() {
		if e := recover(); e != nil {
			log.Println("Error:", e)
		}
	}()
	rows, err := this.db.Query(sqlStr)
	if err != nil {
		log.Println("database select error: ", err.Error())
		return nil
	}
	defer rows.Close()

	var id int
	var title string
	var page int
	var author_id int
	var summary string
	var content string
	var status int
	var created string
	var updated string
	for rows.Next() {
		if err := rows.Scan(&id, &title, &page, &author_id, &summary, &content, &status, &created, &updated); err == nil {
			// log.Println(id, title, page, author_id, summary, content, status, created, updated)
			b := NewBlogByValue(id, title, page, author_id, summary, content, status, created, updated)
			b.readTags()
			bs = append(bs, b)
			// log.Println(len(bs))
		} else {
			log.Println(err.Error())
		}
	}
	return bs
}
func (this *Blog) GetBlogsByTagId(tag_id, start, stop int) (bs []*Blog) {
	this.connectDB()
	defer this.close()
	sqlStr := "select tb_article.id, tb_article.title, tb_article.page, " +
		"tb_article.author_id, tb_article.content, tb_article.summary, tb_article.status," +
		"DATE_FORMAT(created,'%Y年%c月%d日 %T')," +
		" DATE_FORMAT(updated,'%Y年%c月%d日 %T') " +
		"from tb_article,tb_art_tag where tb_article.id = tb_art_tag.art_id " +
		"and tag_id = ?  order by id desc limit " +
		strconv.Itoa(start) + "," + strconv.Itoa(stop)

	defer func() {
		if e := recover(); e != nil {
			log.Println("Error:", e)
		}
	}()
	rows, err := this.db.Query(sqlStr, tag_id)
	if err != nil {
		log.Println("database select error: ", err.Error())
		return nil
	}
	defer rows.Close()

	var id int
	var title string
	var page int
	var author_id int
	var summary string
	var content string
	var status int
	var created string
	var updated string
	for rows.Next() {
		if err := rows.Scan(&id, &title, &page, &author_id, &summary, &content, &status, &created, &updated); err == nil {
			// log.Println(id, title, page, author_id, summary, content, status, created, updated)
			b := NewBlogByValue(id, title, page, author_id, summary, content, status, created, updated)
			b.readTags()
			bs = append(bs, b)
			// log.Println(len(bs))
		} else {
			log.Println(err.Error())
		}
	}
	return bs
	return bs
}
