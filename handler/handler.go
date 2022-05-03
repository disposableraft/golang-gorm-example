package handler

import (
	"lancewakeling.net/go-echo-gorm/article"
)

type Handler struct {
	articleStore article.Store
	// userStore
	// etcStore
	// furthermoreStore
}

/*
Register a new handler.
Its responsibility is basically that of a controller: to bridge requests, data fetching and responses.
*/
func NewHandler(as article.Store) *Handler {
	return &Handler{
		articleStore: as,
		// other stores registered here
	}
}
