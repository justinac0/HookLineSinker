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

	e.GET("/waiting", func(ctx echo.Context) error {
		return renderTemplate(ctx, templates.Waiting())
	})
<<<<<<< Updated upstream
=======

	e.GET("/bite", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "not implemented")
	})

	e.GET("/fight", func(ctx echo.Context) error {
		return renderTemplate(ctx, templates.ArithmeticQuestions())
	})
	e.GET("/battle_fish/:id", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "not implemented")
	})

	e.GET("/result/:fish/:pass", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "not implemented")
	})
>>>>>>> Stashed changes
}
