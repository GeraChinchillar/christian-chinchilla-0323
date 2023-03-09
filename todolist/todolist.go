package todolist

import (
	"chinchilla-christian-0323/entities"
	"fmt"
)

type toDoList struct {
	tasks []entities.Task
}

func NewToDoList(tasks []entities.Task) *toDoList {
	return &toDoList{tasks: tasks}
}

func (t toDoList) AddTask(newTask entities.Task) {
	t.tasks = append(t.tasks, newTask)
}

func (t toDoList) DeleteTask(i int) {
	t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
}

func (t toDoList) PrintTasks() {
	fmt.Println("ID \t Name \t Description \t State \t Begin Date \t Finished Date \t Priority")
	for _, task := range t.tasks {
		task.Print()
	}
}

func (t toDoList) ModifyTask() {

}

func (t toDoList) CountUnfinished() {

}

func (t toDoList) SortPriority() {

}
