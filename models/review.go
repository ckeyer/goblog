package models

import (
// "fmt"
)

type Review struct {
	Id         int64
	Title      string
	Content    string
	Author     string
	Email      string
	ParentType int
	ParentId   int64
}

func NewReview(title, content, author, email string, parenttype int, parentid int64) *Review {
	return &Review{
		Title:      title,
		Content:    content,
		Author:     author,
		Email:      email,
		ParentType: parenttype,
		ParentId:   parentid,
	}
}

func (this *Review) GetArticleReviewsById(articleid int) {

}
