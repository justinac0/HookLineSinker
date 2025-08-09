package handlers

import (
	"HookLineSinker/internal/questions"
	"HookLineSinker/internal/questions/arithmetic"
	"HookLineSinker/web/templates"
	"fmt"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func GenerateArithmeticQuestions() *questions.Grouping {
	qs := questions.NewGrouping("Arithmetic")
	qs.Add(arithmetic.NewArithmeticQuestion("45+37", "82"))
	qs.Add(arithmetic.NewArithmeticQuestion("102-68", "34"))
	qs.Add(arithmetic.NewArithmeticQuestion("14×7", "98"))
	qs.Add(arithmetic.NewArithmeticQuestion("144÷12", "12"))
	qs.Add(arithmetic.NewArithmeticQuestion("58+26", "84"))
	qs.Add(arithmetic.NewArithmeticQuestion("300-175", "125"))
	qs.Add(arithmetic.NewArithmeticQuestion("9×9", "81"))
	qs.Add(arithmetic.NewArithmeticQuestion("120÷8", "15"))
	qs.Add(arithmetic.NewArithmeticQuestion("67+45", "112"))
	qs.Add(arithmetic.NewArithmeticQuestion("500-347", "153"))
	qs.Add(arithmetic.NewArithmeticQuestion("16×5", "80"))
	qs.Add(arithmetic.NewArithmeticQuestion("225÷15", "15"))
	qs.Add(arithmetic.NewArithmeticQuestion("78+99", "177"))
	qs.Add(arithmetic.NewArithmeticQuestion("250-186", "64"))
	qs.Add(arithmetic.NewArithmeticQuestion("13×12", "156"))
	qs.Add(arithmetic.NewArithmeticQuestion("360÷9", "40"))

	return qs
}

func GenerateEnglishQuestions() *questions.Grouping {
	qs := questions.NewGrouping("Gramma")
	return qs
}

func GenerateMusicQuestions() *questions.Grouping {
	qs := questions.NewGrouping("Notes")
	return qs
}

func renderTemplate(ctx echo.Context, component templ.Component) error {
	return component.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func Setup(e *echo.Echo) {
	addition := GenerateArithmeticQuestions()

	e.GET("/", func(ctx echo.Context) error {
		return renderTemplate(ctx, templates.Index())
	})

	e.GET("/casting", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "not implemented")
	})

	e.GET("/waiting", func(ctx echo.Context) error {
		return renderTemplate(ctx, templates.Waiting())
	})

	e.GET("/catch", func(ctx echo.Context) error {
		fmt.Println("catching")
		return renderTemplate(ctx, templates.ArithmeticQuestions(addition))
	})

	e.GET("/fight", func(ctx echo.Context) error {
		return renderTemplate(ctx, templates.ArithmeticQuestions(addition))
	})

	e.GET("/battle_fish/:id", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "not implemented")
	})

	e.POST("/results", func(ctx echo.Context) error {
		formValues, _ := ctx.FormParams()
		passed := true
		for k, v := range formValues {
			if len(v) == 0 {
				passed = false
				break
			}

			if strings.Compare(k, v[0]) != 0 {
				passed = false
				break
			}
		}

		if passed {
			return ctx.String(http.StatusOK, "passed")
		} else {
			return ctx.String(http.StatusOK, "failed")
		}
	})
}
