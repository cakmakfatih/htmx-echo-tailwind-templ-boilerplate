package main

import (
	"echochat/controllers"
	"echochat/internals"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("assets", "./assets")

	internals.InitConfig()

	indexRG := e.Group("/")

	indexController := controllers.NewIndexController(indexRG)
	controllers.RegisterController(indexController)

	e.Logger.Fatal(e.Start(":1323"))
}
