package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Id int `json:"id,omitempty"`
	Title string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	Publisher string `json:"publisher" validate:"required"`
	Year string `json:"year" validate:"required"`
}