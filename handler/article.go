package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler methods are called when setting up a route. Echo passes in the context.
func (h *Handler) GetArticle(c echo.Context) error {
	slug := c.Param("slug")
	a, err := h.articleStore.GetBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "error getting article by slug"})
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"error": err})
	}
	return c.JSON(http.StatusOK, articleResponse(c, a))
}
