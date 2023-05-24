package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Author      Author `json:"author" gorm:"foreignKey:AuthorID;references:ID"`
	AuthorID    uint   `json:"-"`
	Description string `json:"description"`
}
