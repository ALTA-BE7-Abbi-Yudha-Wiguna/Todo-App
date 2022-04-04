package entities

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Nama_Project string `json:"nama_project" form:"nama_project"`
	UserID       uint
	Todo         []Todo
}
