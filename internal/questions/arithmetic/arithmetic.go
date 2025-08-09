package arithmetic

type ArithmeticQuestion struct {
	question string
	answer   int64
}

func NewArithmeticQuestion(name string, answer int64) *ArithmeticQuestion {
	return &ArithmeticQuestion{
		question: name,
		answer:   answer,
	}
}

func (nq *ArithmeticQuestion) GetQuestion() string {
	return nq.question
}

func (nq *ArithmeticQuestion) GradeQuestion() bool {
	return false
}
