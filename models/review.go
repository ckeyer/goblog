package models

import (
"fmt"
)

type Review struct{
Id int64
Title string
Content string
Author string
Email string
ParentType int
ParentId  int64
}

func NewReview(title,content,author,email string){
return &Review{Title:title,Content:content}
