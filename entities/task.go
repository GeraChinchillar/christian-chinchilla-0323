package entities

import (
	"fmt"
	"time"
)

type Task struct {
	id           string
	name         string
	description  string
	state        State
	beginDate    time.Time
	finishedDate time.Time
	priority     Priority
}

func NewTask(id string, name string, description string,
	state State, priority Priority) Task {
	return Task{
		id:           id,
		name:         name,
		description:  description,
		state:        state,
		beginDate:    time.Now(),
		finishedDate: time.Time{},
		priority:     priority,
	}
}

func (t Task) Print() {
	fmt.Println(
		t.id + "\t |" +
			t.name + "\t |" +
			t.description + "\t |" +
			t.state.String() + "\t |" +
			t.beginDate.String() + "\t |" +
			t.finishedDate.String() + "\t |" +
			t.priority.String() + "\t")
}
