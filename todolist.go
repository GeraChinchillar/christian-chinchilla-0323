package main

type toDoList struct {
	tasks []Task
}

func (t toDoList) addTask(newTask Task) {
	t.tasks = append(t.tasks, newTask)
}

func (t toDoList) deleteTask(indice int) {
	t.tasks = append(t.tasks[:indice], t.tasks[indice+1:]...)
}

func (t toDoList) printTasks() {

	for _, task := range t.tasks {
		task.print()
	}
}
