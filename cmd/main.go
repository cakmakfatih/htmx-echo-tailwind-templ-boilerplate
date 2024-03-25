package main

import (
	templates "echochat/templates/layouts"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func HomeHandler(c echo.Context) error {
	return Render(c, http.StatusOK, templates.MainLayout("test"))
}

func main() {
	e := echo.New()
	e.Static("assets", "./public")

	e.GET("/", HomeHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
