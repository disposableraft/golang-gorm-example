package handler

import "github.com/labstack/echo/v4"

func (h *Handler) Register(e *echo.Echo) {
	e.GET("/:slug", h.GetArticle)
}
