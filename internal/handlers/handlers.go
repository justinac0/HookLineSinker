package handlers

import (
	"HookLineSinker/internal/common"
	"HookLineSinker/internal/questions"
	"HookLineSinker/web/templates"
	"HookLineSinker/web/templates/components"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func getMathQuestions() *questions.QuestionGrouping {
	qs := questions.NewQuestionGrouping("Math")

	qs.Add(questions.NewSingleQuestion("6 x 6 ", "36"))

	qs.Add(questions.NewSingleQuestion("81 ÷ 9 ", "9"))

	qs.Add(questions.NewSingleQuestion("87 + 6 ", "93"))

	qs.Add(questions.NewSingleQuestion("90 - 18 ", "72"))

	qs.Add(questions.NewSingleQuestion("80 - 27 ", "53"))

	qs.Add(questions.NewSingleQuestion("70 + 2 ", "72"))

	qs.Add(questions.NewSingleQuestion("90 ÷ 6 ", "15"))

	qs.Add(questions.NewSingleQuestion("45 ÷ 9 ", "5"))

	qs.Add(questions.NewSingleQuestion("50 ÷ 5 ", "10"))

	qs.Add(questions.NewSingleQuestion("67 + 17 ", "84"))

	qs.Add(questions.NewSingleQuestion("9 x 8 ", "72"))

	qs.Add(questions.NewSingleQuestion("56 ÷ 7 ", "8"))

	qs.Add(questions.NewSingleQuestion("64 ÷ 8 ", "8"))

	qs.Add(questions.NewSingleQuestion("56 - 19 ", "37"))

	qs.Add(questions.NewSingleQuestion("95 + 7 ", "102"))
	return qs

}

func getEnglishQuestions() *questions.QuestionGrouping {
	qs := questions.NewQuestionGrouping("English")

	qs.Add(questions.NewMultipleChoiceQuestion("What is the plural of “mouse”?", []string{
		"Mouses", "Mouse", "Mice",
	}, "C"))

	qs.Add(questions.NewMultipleChoiceQuestion("Which of the following word is a verb?", []string{
		"Walk", "Pink", "Water Bottle",
	}, "A"))

	qs.Add(questions.NewMultipleChoiceQuestion("What is the opposite of the word “day”?", []string{
		"Blue", "Night", "Days",
	}, "B"))

	qs.Add(questions.NewMultipleChoiceQuestion("How many syllables are in the word “banana”", []string{
		"1", "2", "3",
	}, "C"))

	qs.Add(questions.NewMultipleChoiceQuestion("How many syllables are in the word “musical”", []string{
		"1", "2", "3",
	}, "C"))

	qs.Add(questions.NewMultipleChoiceQuestion("What is the plural of “child”", []string{
		"Childs", "Children", "Child",
	}, "B"))

	qs.Add(questions.NewMultipleChoiceQuestion("How many syllables are in the word “statement”", []string{
		"1", "2", "3",
	}, "B"))

	qs.Add(questions.NewMultipleChoiceQuestion("How many syllables are in the word “bottle”", []string{
		"1", "2", "3",
	}, "B"))

	qs.Add(questions.NewMultipleChoiceQuestion("Which of the following is the correct spelling:", []string{
		"Neccessary", "Necessary", " Necassory",
	}, "B"))

	qs.Add(questions.NewMultipleChoiceQuestion("Which of the following is the correct spelling:", []string{
		"Believe", "Beleive", " Bilieve",
	}, "A"))

	qs.Add(questions.NewMultipleChoiceQuestion("Which of the following is the correct spelling:", []string{
		"Broccoli", "Brocoli", " Broccolli",
	}, "A"))

	qs.Add(questions.NewMultipleChoiceQuestion("How many syllables are in the word “Broccoli”:", []string{
		"1", "2", "3",
	}, "A"))

	qs.Add(questions.NewMultipleChoiceQuestion("How many syllables are in the word “Seperate”:", []string{
		"1", "2", "3",
	}, "C"))

	qs.Add(questions.NewMultipleChoiceQuestion("Which of the following is a verb", []string{
		"Direction", "Swim", "Fried Chicken",
	}, "B"))

	qs.Add(questions.NewMultipleChoiceQuestion("How many syllables are in the word “Directory”:", []string{
		"2", "3", "4",
	}, "C"))

	qs.Add(questions.NewMultipleChoiceQuestion("How many syllables are in the word “Mackerel”:", []string{
		"2", "3", "4",
	}, "B"))

	qs.Add(questions.NewMultipleChoiceQuestion("What is the synonym “Clever”:", []string{
		"Intelligent", "Interesting", "Clown",
	}, "A"))

	qs.Add(questions.NewMultipleChoiceQuestion("Which of the following word is a noun", []string{
		"Lion", "Run", "Route",
	}, "A"))

	qs.Add(questions.NewMultipleChoiceQuestion("Which of the following word is a noun", []string{
		"Table", "Sing", "Fast",
	}, "A"))

	return qs
}

func getMusicQuestions() *questions.QuestionGrouping {
	qs := questions.NewQuestionGrouping("Music")

	qs.Add(questions.NewSingleQuestion("eng", "ong"))
	qs.Add(questions.NewMultipleChoiceQuestion("How many strings does a standard guitar have?", []string{
		"6", "7", "8",
	}, "A"))

	qs.Add(questions.NewMultipleChoiceQuestion("Which of these is the highest singing voice?", []string{
		"Soprano", "Bass", "Alto",
	}, "A"))

	qs.Add(questions.NewMultipleChoiceQuestion("What do musicians read to play music?", []string{
		"Map", "Music notes", "Recipe",
	}, "B"))

	qs.Add(questions.NewMultipleChoiceQuestion("Which family of instruments does the trumpet belong to?", []string{
		"String", "Brass", "Percussion",
	}, "B"))

	qs.Add(questions.NewMultipleChoiceQuestion("Which instrument do you blow air into to play?", []string{
		"Violin", "Flute", "Guitar",
	}, "B"))

	qs.Add(questions.NewMultipleChoiceQuestion("What do you call the speed of the music?", []string{
		"Tempo", "Volume", "Melody",
	}, "A"))

	qs.Add(questions.NewMultipleChoiceQuestion("How many beats does a whole note usually get?", []string{
		"2", "4", "8",
	}, "B"))

	qs.Add(questions.NewMultipleChoiceQuestion("Which instrument has strings and is played with a bow?", []string{
		"Flute", "Drum", "Violin",
	}, "C"))

	qs.Add(questions.NewMultipleChoiceQuestion("What do we call a group of people singing together?", []string{
		"Band", "Choir", "Orchestra",
	}, "B"))

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
		index := rand.Int() % int(common.FishCount)
		fishType := common.FishType(index)
		return renderTemplate(ctx, components.Fish(fishType))
	})

	e.GET("/fight/:id", func(ctx echo.Context) error {
		varID := ctx.QueryParams().Get("id")
		num, err := strconv.Atoi(varID)
		if err != nil {
			return renderTemplate(ctx, templates.Index())
		}

		fishType := common.FishType(num)
		var group *questions.QuestionGrouping
		switch fishType {
		case common.EnglishFish:
			group = getEnglishQuestions()
		case common.MathFish:
			group = getMathQuestions()
		case common.MusicFish:
			group = getMusicQuestions()
		}

		return renderTemplate(ctx, templates.Quiz(group.GetTitle(), group.GetRandomQuestion()))
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
