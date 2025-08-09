package questions

type Question interface {
	GetQuestion() string
	GradeQuestion() bool
}

type Grouping struct {
	title     string
	questions []Question
}

func NewGrouping(title string) *Grouping {
	return &Grouping{
		title: title,
	}
}

func (g *Grouping) Add(q Question) {
	g.questions = append(g.questions, q)
}

func (g *Grouping) Print() {

}
