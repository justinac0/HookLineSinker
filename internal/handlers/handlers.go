package handlers

import (
	"HookLineSinker/internal/questions"
	"HookLineSinker/internal/questions/arithmetic"
	"HookLineSinker/web/templates"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func GenerateArithmeticQuestions() *questions.Grouping {
	arithmeticQuestions := questions.NewGrouping("Addition")
	arithmeticQuestions.Add(arithmetic.NewArithmeticQuestion("12+4", 16))
	arithmeticQuestions.Add(arithmetic.NewArithmeticQuestion("45+27", 72))
	arithmeticQuestions.Add(arithmetic.NewArithmeticQuestion("103+89", 192))

	return arithmeticQuestions
}

func renderTemplate(ctx echo.Context, component templ.Component) error {
	return component.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func Setup(e *echo.Echo) {
	e.GET("/", func(ctx echo.Context) error {
		return renderTemplate(ctx, templates.Index())
	})

	e.GET("/casting", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "not implemented")
	})

	e.GET("/waiting", func(ctx echo.Context) error {
		return renderTemplate(ctx, templates.Waiting())
	})

	e.GET("/bite", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "not implemented")
	})

	addition := GenerateArithmeticQuestions()
	fmt.Println("addition questions:", addition)

	e.GET("/fight", func(ctx echo.Context) error {
		return renderTemplate(ctx, templates.ArithmeticQuestions())
	})

	e.GET("/battle_fish/:id", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "not implemented")
	})

	e.GET("/result/:fish/:pass", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "not implemented")
	})
}
