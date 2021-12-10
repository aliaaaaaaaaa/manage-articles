package model

import "github.com/jinzhu/gorm"

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  int    `json:"user_id,omitempty"`
	Author  string `json:"author"`
	Tags    []Tage `json:"tags"`
	gorm.Model
}
