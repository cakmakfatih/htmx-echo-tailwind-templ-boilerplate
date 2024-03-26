package main

import (
	"echochat/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("assets", "./assets")

	indexRG := e.Group("/")

	indexController := controllers.NewIndexController(indexRG)
	controllers.RegisterController(indexController)

	e.Logger.Fatal(e.Start(":1323"))
}
