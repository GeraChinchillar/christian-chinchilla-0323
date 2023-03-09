package entities

type State interface {
	SetStateId(id int)
	SetStateName(name string)
	String() string
}

type Pendant struct {
	id   int
	name string
}

func (p Pendant) SetStateId(id int) {
	p.id = id
}

func (p Pendant) SetStateName(name string) {
	p.name = name
}

func (p Pendant) String() string {
	return string(p.id) + " = " + p.name
}
