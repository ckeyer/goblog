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

type Blog struct {
	ID          int
	Title       string
	Page        int
	AuthorID    int
	Content     string
	Status      int
	CreatedTime string
	UpdateTme   string

	db *sql.DB
}

func NewBlog() (b *Blog) {
	b = &Blog{}
	return b
}
func NewBlogByValue(id int, title string, page int, author_id int, content string, status int, created string, updated string) (b *Blog) {
	b = &Blog{
		ID:          id,
		Title:       title,
		Page:        page,
		AuthorID:    author_id,
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

func (this *Blog) Read() {
	this.connectDB()
	defer this.close()
	sqlStr := "select id, title, page, author_id, content, status, DATE_FORMAT(created,'%Y年%c月%d日'), DATE_FORMAT(updated,'%Y年%c月%d日') from tb_article "

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
	var content string
	var status int
	var created string
	var updated string

	for rows.Next() {
		if err := rows.Scan(&id, &title, &page, &author_id, &content, &status, &created, &updated); err == nil {
			log.Println(id, title, page, author_id, content, status, created, updated)
		} else {
			log.Println(err.Error())
		}
	}
}
func (this *Blog) GetBlogList(start, stop int) (bs *list.List) {
	this.connectDB()
	defer this.close()
	sqlStr := "select id, title, page, author_id, content, status, DATE_FORMAT(created,'%Y年%c月%d日'), DATE_FORMAT(updated,'%Y年%c月%d日') from tb_article limit " +
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
	var content string
	var status int
	var created string
	var updated string

	bs = list.New()
	for rows.Next() {
		if err := rows.Scan(&id, &title, &page, &author_id, &content, &status, &created, &updated); err == nil {
			log.Println(id, title, page, author_id, content, status, created, updated)
			bs.PushBack(NewBlogByValue(id, title, page, author_id, content, status, created, updated))
		} else {
			log.Println(err.Error())
		}
	}
	return bs
}
func (this *Blog) ToMap() map[string]string {
	bm := make(map[string]string)
	bm["id"] = strconv.Itoa(this.ID)
	bm["title"] = this.Title
	bm["page"] = strconv.Itoa(this.Page)
	bm["author_id"] = strconv.Itoa(this.AuthorID)
	bm["content"] = this.Content
	bm["status"] = strconv.Itoa(this.Status)
	bm["created"] = this.CreatedTime
	bm["updated"] = this.UpdateTme
	return bm
}
func (this *Blog) ReadBlogByID(id_ string) error {
	this.connectDB()
	defer this.close()
	sqlStr := "select id, title, page, author_id, content, status, DATE_FORMAT(created,'%Y年%c月%d日'), DATE_FORMAT(updated,'%Y年%c月%d日') from tb_article "

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
	var content string
	var status int
	var created string
	var updated string

	if err := rows.Scan(&id, &title, &page, &author_id, &content, &status, &created, &updated); err == nil {
		log.Println(id, title, page, author_id, content, status, created, updated)
		this.ID = id
		this.Title = title
		this.Page = page
		this.AuthorID = author_id
		this.Content = content
		this.Status = status
		this.CreatedTime = created
		this.UpdateTme = updated
	} else {
		return err
	}
	return nil
}
