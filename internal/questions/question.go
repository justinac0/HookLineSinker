package questions

import "math/rand"

type IQuestion interface {
	GetQuestions() []string
	GetQuestionCount() int
	GetAnswer() string
	GetDescription() string
}

type QuestionGrouping struct {
	title        string
	questionPool []IQuestion
}

type Question struct {
	description string
	questions   []string
	answer      string
}

func (q *Question) GetQuestions() []string {
	return q.questions
}

func (q *Question) GetQuestionCount() int {
	return len(q.questions)
}

func (q *Question) GetAnswer() string {
	return q.answer
}

func (q *Question) GetDescription() string {
	return q.description
}

func NewMultipleChoiceQuestion(description string, questions []string, answer string) *Question {
	return &Question{
		description,
		questions,
		answer,
	}
}

func NewSingleQuestion(question string, answer string) *Question {
	return &Question{
		description: "",
		questions:   []string{question},
		answer:      answer,
	}
}

func NewQuestionGrouping(title string) *QuestionGrouping {
	return &QuestionGrouping{
		title:        title,
		questionPool: make([]IQuestion, 0),
	}
}

func (g *QuestionGrouping) Add(q IQuestion) {
	if g.questionPool == nil {
		g.questionPool = make([]IQuestion, 0)
	}
	g.questionPool = append(g.questionPool, q)
}

func (g *QuestionGrouping) GetTitle() string {
	return g.title
}

func (g *QuestionGrouping) GetQuestionPool() []IQuestion {
	if g.GetQuestionCount() == 0 {
		return nil
	}
	return g.questionPool
}

func (g *QuestionGrouping) GetRandomQuestion() IQuestion {
	index := rand.Int() % g.GetQuestionCount()
	return g.GetQuestionPool()[index]
}

func (g *QuestionGrouping) GetQuestionCount() int {
	return len(g.questionPool)
}
