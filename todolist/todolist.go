package todolist

import "chinchilla-christian-0323/entities"

type toDoList struct {
	tasks []entities.Task
}

func NewToDoList(tasks []entities.Task) *toDoList {
	return &toDoList{tasks: tasks}
}

func (t toDoList) AddTask(newTask entities.Task) {
	t.tasks = append(t.tasks, newTask)
}

func (t toDoList) DeleteTask(indice int) {
	t.tasks = append(t.tasks[:indice], t.tasks[indice+1:]...)
}

func (t toDoList) PrintTasks() {

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
