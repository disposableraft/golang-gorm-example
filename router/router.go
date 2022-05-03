package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

/*
router sets up an instance of echo server and applies any middleware, in this case logging.
*/
func New() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Use(middleware.Logger())
	return e
}
