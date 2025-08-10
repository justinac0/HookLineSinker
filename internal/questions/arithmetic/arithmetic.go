package arithmetic

type ArithmeticQuestion struct {
	questions []string
	answer    string
}

func NewArithmeticQuestion(name string, answer string) *ArithmeticQuestion {
	return &ArithmeticQuestion{
		questions: []string{name},
		answer:    answer,
	}
}

func (nq *ArithmeticQuestion) GetQuestion() string {

	return nq.questions[0]
}

func (nq *ArithmeticQuestion) GetAnswer() string {
	return nq.answer
}

func (nq *ArithmeticQuestion) GetQuestionCount() int {
	return len(nq.questions)
}
