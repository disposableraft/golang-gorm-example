package model

import "gorm.io/gorm"

/*
The model represents a table's fields in the db. That's all it does!
*/
type Article struct {
	gorm.Model
	Slug  string
	Title string
	Body  string
	// ...
}

/*
If a model struct doesn't match the table's name it can be overridden using the Tabler interface.
*/
func (Article) TableName() string {
	return "content"
}
