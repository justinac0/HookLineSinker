package handlers

import (
	"HookLineSinker/internal/common"
	"HookLineSinker/internal/questions"
	"HookLineSinker/web/templates"
	"HookLineSinker/web/templates/components"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func GenerateArithmeticQuestions() *questions.QuestionGrouping {
	qs := questions.NewQuestionGrouping("Arithmetic")
	qs.Add(questions.NewMultipleChoiceQuestion("Which statement is correct?", []string{"1+1", "test"}, "A"))

	return qs
}

func GenerateEnglishQuestions() *questions.QuestionGrouping {
	qs := questions.NewQuestionGrouping("Gramma")
	return qs
}

func GenerateMusicQuestions() *questions.QuestionGrouping {
	qs := questions.NewQuestionGrouping("Notes")
	return qs
}

func renderTemplate(ctx echo.Context, component templ.Component) error {
	return component.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func generateMOTDList() []string {
	var motdList []string
	motdList = append(motdList, "Whale, whale, whale… the whale shark is the largest fish!")
	motdList = append(motdList, "Holy mackerel! Some fish can glow in the dark!")
	motdList = append(motdList, "Shrimply amazing! Some shrimps can see more colors than humans :o")
	motdList = append(motdList, "Cod you believe it? Sharks have no bones (except for their teeth)!")
	motdList = append(motdList, "I bait you did not know that some fish glow in the dark!")
	motdList = append(motdList, "Major pikes in productivity!")
	motdList = append(motdList, "You're krilling it.")
	motdList = append(motdList, "You are off the scale.")
	motdList = append(motdList, "Are you fin-ished playing?")
	motdList = append(motdList, "Now this is reel education :>D")
	return motdList
}

func Setup(e *echo.Echo) {
	addition := GenerateArithmeticQuestions()
	motdList := generateMOTDList()

	e.GET("/motd", func(ctx echo.Context) error {
		index := rand.Int() % len(motdList)
		return ctx.String(http.StatusOK, motdList[index])
	})

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
		index := rand.Int() % int(common.FishCount-1)
		fishType := common.FishType(index)
		return renderTemplate(ctx, components.Fish(fishType))
	})

	// questions := addition.GetQuestions()
	e.GET("/fight", func(ctx echo.Context) error {
		question := addition.GetRandomQuestion()

		log.Println(question)

		return renderTemplate(ctx, templates.Quiz(addition.GetTitle(), question))
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

		return renderTemplate(ctx, templates.Results(passed))
	})
}
