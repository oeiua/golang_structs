package zoo

type Leg struct {
	Side     string // left or right
	Position string // top or bottom
}

type Head struct {
	fangs bool
}

type Animal struct {
	Head
	Name string
	Type string
	Legs []Leg
}

type Cage struct {
	Animals []Animal
}

type Zoo struct {
	Cages []*Cage
}

func (animal Animal) MakePredator() {
	animal.fangs = true
}
func (animal Animal) MakeHerbivor() {
	animal.fangs = false
}

func (animal Animal) IsPredator() bool {
	if animal.fangs {
		return true
	} else {
		return false
	}
}
