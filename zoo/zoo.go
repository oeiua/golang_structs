package zoo

type Leg struct {
	Side     string // left or right
	Position string // top or bottom
}

type Head struct {
	Fangs bool
}

type Animal struct {
	Name string
	Type string
	Legs []Leg
	Head Head
}

type Cage struct {
	Animals []Animal
}

type Zoo struct {
	Cages []*Cage
}
