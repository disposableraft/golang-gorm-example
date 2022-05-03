package handler

import (
	"github.com/labstack/echo/v4"
	"lancewakeling.net/go-echo-gorm/model"
)

type singleArticleResponse struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func articleResponse(c echo.Context, a *model.Article) *singleArticleResponse {
	return &singleArticleResponse{
		Slug:  a.Slug,
		Title: a.Title,
		Body:  a.Body,
	}
}
