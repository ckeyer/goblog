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
}

func NewReview(title,content,author,email string){
}
