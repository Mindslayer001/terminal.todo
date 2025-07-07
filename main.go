package main

import (
	"bufio"
	"fmt"
	"os"

	tem "github.com/mindslayer001/terminal.todo/Todo"
)

func main() {
	var T tem.TODOList
	T.LoadTODO()
	for {
		T.PrintTODO()
		printChoices()
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			var Title string
			fmt.Println("Enter the new todo Item")
			reader := bufio.NewScanner(os.Stdin)
			if reader.Scan() {
				Title = reader.Text()
				T.CreateTODO(Title, false)
			}
		case 2:
			var index int
			fmt.Println("Enter the todo Number")
			fmt.Scanln(&index)
			T.Complete(index)
		case 3:
			var index int
			fmt.Println("Enter the todo Number")
			fmt.Scanln(&index)
			T.DeleteTodo(index)
		case 4:
			T.CheckTODONodes()
			return
		default:
			fmt.Println("are u retared please Choose the correct from below!:")
			printChoices()
		}
	}
}

func printChoices() {
	println("Choose from below options")
	println("1 Create New todo")
	println("2 Change the Status of a certain todo")
	println("3 Delete the todo")
	println("4 Exit")
}
