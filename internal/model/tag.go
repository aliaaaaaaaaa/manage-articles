package model

import "github.com/jinzhu/gorm"

type Tage struct {
	Name string `json:"name"`
	gorm.Model
}
