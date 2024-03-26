package main

import (
	"echochat/controllers"
	"echochat/internals"
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	e := echo.New()
	e.Static("assets", "./assets")

	internals.InitConfig()

	indexRG := e.Group("/")

	indexController := controllers.NewIndexController(indexRG)
	controllers.RegisterController(indexController)

	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%v", os.Getenv("PORT"))))
}
