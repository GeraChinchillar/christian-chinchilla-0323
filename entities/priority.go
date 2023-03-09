package entities

type Priority interface {
	SetPriorityId(id int)
	SetPriorityName(name string)
	String() string
}

type Low struct {
	id   int
	name string
}

func (l Low) SetPriorityId(id int) {
	l.id = id
}

func (l Low) SetPriorityName(name string) {
	l.name = name
}

func (l Low) String() string {
	return string(l.id) + " = " + l.name
}
