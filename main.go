package main

import (
	"chinchilla-christian-0323/entities"
	"chinchilla-christian-0323/todolist"
)

func main() {

	pendantState := entities.Pendant{}
	pendantState.SetStateId(1)
	pendantState.SetStateName("Pendant")

	lowPriority := entities.Low{}
	lowPriority.SetPriorityId(1)
	lowPriority.SetPriorityName("Pendant")

	tsk1 := entities.NewTask("1", "Task1", "Task1", pendantState, lowPriority)
	tsk2 := entities.NewTask("2", "Task1", "Task1", pendantState, lowPriority)
	tsk3 := entities.NewTask("3", "Task1", "Task1", pendantState, lowPriority)
	tsk4 := entities.NewTask("4", "Task1", "Task1", pendantState, lowPriority)

	newToDoList := todolist.NewToDoList([]entities.Task{tsk1, tsk2, tsk3, tsk4})
	newToDoList.PrintTasks()

}
