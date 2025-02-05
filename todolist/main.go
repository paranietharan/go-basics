package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	taskName   string
	isFinished bool
}

func newTask(taskName string) *Task {
	newTask := Task{
		taskName:   taskName,
		isFinished: false,
	}

	return &newTask
}

// array 2 store the tasks
var taskList []Task

func addTask(taskName string) {
	taskList = append(taskList, *newTask(taskName))
	fmt.Println("Task added: " + taskName)
}

func deleteTask(index int) {
	if index < 0 || index > len(taskList) {
		fmt.Println("You entered invalide Index")
	}

	taskList = append(taskList[:index], taskList[index+1:]...)
	fmt.Println("Task deleted successfully!")
}

func viewTask() {
	if len(taskList) == 0 {
		fmt.Println("Task is empty!")
	}
	fmt.Println("Tasks : ")
	for i, task := range taskList {
		status := "[\u2718]"
		if task.isFinished {
			status = "[\u2714]"
		}
		fmt.Printf("%d: %s %s\n", i+1, status, task.taskName)
	}
}

func markTaskComplete(index int) {
	if index < 0 || index > len(taskList) {
		fmt.Println("You entered invalide Index")
	}
	taskList[index].isFinished = true
	fmt.Println("Task marked as complete!")
}

func main() {
	var choice int
	reader := bufio.NewReader(os.Stdin)

	// infinite loop
	i := 1
	for i > 0 {
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task as Complete")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		reader.ReadString('\n')

		switch choice {
		case 1:
			//var taskName string
			// fmt.Scanf("%s\n", &taskName)
			// addTask(taskName)
			fmt.Print("Enter the task name : ")
			taskName, _ := reader.ReadString('\n') // Waits until Enter is pressed
			taskName = taskName[:len(taskName)-1]  // Remove newline character
			addTask(taskName)
		case 2:
			viewTask()
		case 3:
			fmt.Print("Enter task number to mark as complete: ")
			var index int
			fmt.Scan(&index)
			markTaskComplete(index - 1)
		case 4:
			fmt.Print("Enter task number to mark as complete: ")
			var index int
			fmt.Scan(&index)
			deleteTask(index - 1)
		case 5:
			fmt.Println("Exiting...")
			i = 0 // to exit the prgram initialize 0 to the i
		}
	}
}
