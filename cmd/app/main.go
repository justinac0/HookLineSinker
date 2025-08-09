package main

import (
	"HookLineSinker/internal/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("static/", "web/static/")

	handlers.Setup(e)

	e.HideBanner = true
	e.Logger.Fatal(e.Start("127.0.0.1:3000"))
}
