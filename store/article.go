package store

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	"lancewakeling.net/go-echo-gorm/model"
)

/*
All database logic is encapsilated within stores
*/
type ArticleStore struct {
	db *gorm.DB
}

func NewArticleStore(db *gorm.DB) *ArticleStore {
	return &ArticleStore{
		db: db,
	}
}

/*
An example of store logic. This is close to a "model" in django or rails.
*/
func (as *ArticleStore) GetBySlug(s string) (*model.Article, error) {
	a := &model.Article{Slug: s}
	err := as.db.Where("slug = ?", s).First(&a).Error
	if err != nil {
		fmt.Println("article not found by slug")
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, err
		}
		return nil, err
	}
	return a, nil
}
