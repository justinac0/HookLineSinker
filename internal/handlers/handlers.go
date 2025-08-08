package handlers

import (
	"HookLineSinker/web/templates"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func renderTemplate(ctx echo.Context, component templ.Component) error {
	return component.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func Setup(e *echo.Echo) {
	e.GET("/", func(ctx echo.Context) error {
		return renderTemplate(ctx, templates.Index())
	})
}
