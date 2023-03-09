package main

import "time"

func main() {

	task1 := Task{
		id:           "1",
		name:         "",
		description:  "",
		state:        nil,
		beginDate:    time.Time{},
		finishedDate: time.Time{},
		priority:     Priority{},
	}

	task2 := Task{
		id:           "2",
		name:         "",
		description:  "",
		state:        nil,
		beginDate:    time.Time{},
		finishedDate: time.Time{},
		priority:     Priority{},
	}
	task3 := Task{
		id:           "3",
		name:         "",
		description:  "",
		state:        nil,
		beginDate:    time.Time{},
		finishedDate: time.Time{},
		priority:     Priority{},
	}
	task4 := Task{
		id:           "4",
		name:         "",
		description:  "",
		state:        nil,
		beginDate:    time.Time{},
		finishedDate: time.Time{},
		priority:     Priority{},
	}

	newToDoList := toDoList{tasks: []Task{task1, task2, task3, task4}}
}
