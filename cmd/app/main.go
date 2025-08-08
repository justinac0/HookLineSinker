package main

import (
	"HookLineSinker/internal/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("static/css", "web/static/css")
	e.Static("static/img", "web/static/img")
	e.Static("static/js", "web/static/js")

	handlers.Setup(e)

	e.HideBanner = true
	e.Logger.Fatal(e.Start(":3000"))
}
