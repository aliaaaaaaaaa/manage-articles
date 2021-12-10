package model

import "github.com/jinzhu/gorm"

type Article struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	UserID int    `json:"user_id,omitempty"`
	Tags   []Tage `json:"tags"`
	gorm.Model
}
