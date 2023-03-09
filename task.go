package main

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

func (t Task) print() {
	fmt.Println(
		t.id + "\t" +
			t.name + "\t" +
			t.description + "\t" +
			t.state.String() + "\t" +
			t.id + "\t" +
			t.id + "\t" +
			t.id + "\t")
}