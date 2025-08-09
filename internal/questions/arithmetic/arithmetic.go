package arithmetic

type ArithmeticQuestion struct {
	question string
	answer   string
}

func NewArithmeticQuestion(name string, answer string) *ArithmeticQuestion {
	return &ArithmeticQuestion{
		question: name,
		answer:   answer,
	}
}

func (nq *ArithmeticQuestion) GetQuestion() string {
	return nq.question
}

func (nq *ArithmeticQuestion) GetAnswer() string {
	return nq.answer
}

func (nq *ArithmeticQuestion) GradeQuestion() bool {
	return false
}
