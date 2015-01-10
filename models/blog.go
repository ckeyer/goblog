package models

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"log"
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

	dbStr := beego.AppConfig.String("sql_db")
	connStr := beego.AppConfig.String("sql_conn_str")
	var err error
	b.db, err = sql.Open(dbStr, connStr)

	if err != nil {
		log.Println("database initialize error : ", err.Error())
		// b.db = db
	}
	return b
}
func (b *Blog) Read() {
	sqlStr := "select id, title, page, author_id, content, status, created, updated from tb_article limit 0,5"
	if err := b.db.Ping(); err != nil {
		log.Println("database connect error: ", err.Error())
	}
	rows, err := b.db.Query(sqlStr)
	if err != nil {
		log.Println("database select error: ", err.Error())
		return
	} else {
		log.Println("sdfsfasdfadsf")
	}
	defer rows.Close()
	// cols, _ := rows.Columns()

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
func (b *Blog) Close() {
	if b.db.Ping() == nil {
		b.db.Close()
	}
}
