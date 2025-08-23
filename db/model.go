package db

import "gorm.io/gorm"

type Article struct {
	*gorm.Model
	Author      string  `json:"author" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Likes       float64 `json:"likes"`
}
