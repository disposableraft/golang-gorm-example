package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"lancewakeling.net/go-echo-gorm/model"
)

/*
connect to the database
*/
func New() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})

	if !db.Migrator().HasTable(&model.Article{}) {
		db.AutoMigrate(&model.Article{})

		articles := []model.Article{
			{
				Slug:  "hello-world",
				Title: "Hello World",
				Body:  "Hello Beautiful World!",
			},
		}

		db.Save(&articles)
	}

	if err != nil {
		panic(fmt.Sprintf("failed to connect to database %s", err))
	}

	fmt.Println("Connected to database")
	return db
}
