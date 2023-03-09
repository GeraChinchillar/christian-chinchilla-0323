package todolist

import (
	"chinchilla-christian-0323/entities"
	"reflect"
	"testing"
)

func TestNewToDoList(t *testing.T) {
	type args struct {
		tasks []entities.Task
	}

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

	tasks := []entities.Task{tsk1, tsk2, tsk3, tsk4}

	tests := []struct {
		name string
		args args
		want *toDoList
	}{
		{
			name: "Test NewToDoList",
			args: args{tasks: tasks},
			want: &toDoList{tasks: tasks},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewToDoList(tt.args.tasks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewToDoList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toDoList_AddTask(t1 *testing.T) {
	type fields struct {
		tasks []entities.Task
	}
	type args struct {
		newTask entities.Task
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := toDoList{
				tasks: tt.fields.tasks,
			}
			t.AddTask(tt.args.newTask)
		})
	}
}

func Test_toDoList_CountUnfinished(t1 *testing.T) {
	type fields struct {
		tasks []entities.Task
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := toDoList{
				tasks: tt.fields.tasks,
			}
			t.CountUnfinished()
		})
	}
}

func Test_toDoList_DeleteTask(t1 *testing.T) {
	type fields struct {
		tasks []entities.Task
	}
	type args struct {
		indice int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := toDoList{
				tasks: tt.fields.tasks,
			}
			t.DeleteTask(tt.args.indice)
		})
	}
}

func Test_toDoList_ModifyTask(t1 *testing.T) {
	type fields struct {
		tasks []entities.Task
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := toDoList{
				tasks: tt.fields.tasks,
			}
			t.ModifyTask()
		})
	}
}

func Test_toDoList_PrintTasks(t1 *testing.T) {
	type fields struct {
		tasks []entities.Task
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := toDoList{
				tasks: tt.fields.tasks,
			}
			t.PrintTasks()
		})
	}
}

func Test_toDoList_SortPriority(t1 *testing.T) {
	type fields struct {
		tasks []entities.Task
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := toDoList{
				tasks: tt.fields.tasks,
			}
			t.SortPriority()
		})
	}
}
