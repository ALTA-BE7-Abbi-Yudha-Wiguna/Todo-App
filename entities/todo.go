package entities

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Todo     string `json:"todo" form:"todo"`
	Status   string `json:"status" form:"status"`
	UserID   uint
	Projects []Project
}
