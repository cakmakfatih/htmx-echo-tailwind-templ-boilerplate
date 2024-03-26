package controllers

import (
	"echochat/internals"
	pages "echochat/templates/pages"
	"github.com/labstack/echo/v4"
	"net/http"
)

type IndexController struct {
	rg *echo.Group
}

func indexHandler(c echo.Context) error {
	return internals.RenderTempl(c, http.StatusOK, pages.Home())
}

func (c *IndexController) Register() {
	c.rg.GET("", indexHandler)
}

func NewIndexController(rg *echo.Group) *IndexController {
	return &IndexController{
		rg: rg,
	}
}
