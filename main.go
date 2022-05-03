package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"lancewakeling.net/go-echo-gorm/db"
	"lancewakeling.net/go-echo-gorm/handler"
	"lancewakeling.net/go-echo-gorm/router"
	"lancewakeling.net/go-echo-gorm/store"
)

/*
main is called and runs until it crashes or is interupted.
*/
func main() {
	// router sets up echo, applies middle ware and returns Echo
	r := router.New()

	// Connect to the db with gorm
	db := db.New()

	// /Stores contain all logic relating to data retrieval
	articleStore := store.NewArticleStore(db)

	// The handler is a collection of stores. The handler also has methods, such as h.GetArticle or h.CreateNewArticle, etc.
	h := handler.NewHandler(articleStore)

	// Registering the handler defines the endpoints and connects them to handler methods.
	// For example h.GetArticle is the glue that connects the request, the article store
	// and the response for GET "/:slug"
	h.Register(r)

	// The alpha and omega
	r.Logger.Fatal(r.Start(":9999"))
}
