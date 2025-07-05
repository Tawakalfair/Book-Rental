package book

import "gorm.io/gorm"

type Book struct {
	gorm.Model        // Include ID, CreatedAt, UpdatedAt, DeletedAt
	Title      string `json:"title"`
	Author     string `json:"author"`
}
