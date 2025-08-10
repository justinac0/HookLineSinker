package questions

import "math/rand"

type IQuestion interface {
	GetQuestions() []string
	GetQuestionCount() int
	GetAnswer() string
}

type Grouping struct {
	title        string
	questionPool []IQuestion
}

type Question struct {
	questions []string
	answer    string
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

func NewQuestion(questions []string, answer string) *Question {
	return &Question{
		questions,
		answer,
	}
}

func NewGrouping(title string) *Grouping {
	return &Grouping{
		title:        title,
		questionPool: make([]IQuestion, 0),
	}
}

func (g *Grouping) Add(q IQuestion) {
	if g.questionPool == nil {
		g.questionPool = make([]IQuestion, 0)
	}
	g.questionPool = append(g.questionPool, q)
}

func (g *Grouping) GetTitle() string {
	return g.title
}

func (g *Grouping) GetQuestionPool() []IQuestion {
	if g.GetQuestionCount() == 0 {
		return nil
	}
	return g.questionPool
}

func (g *Grouping) GetRandomQuestion() IQuestion {
	index := rand.Int() % g.GetQuestionCount()
	return g.GetQuestionPool()[index]
}

func (g *Grouping) GetQuestionCount() int {
	return len(g.questionPool)
}
