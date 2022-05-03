package article

import "lancewakeling.net/go-echo-gorm/model"

type Store interface {
	GetBySlug(string) (*model.Article, error)
}
