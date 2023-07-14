package main

import (
	"fmt"
	"todo-golang/controller"
	"todo-golang/controller2"
)

func DisplayInterface() int {
	fmt.Println(("Choose no. based on operation you want to perform..."))
	fmt.Print("1. To add todo\n2. To get a todo\n3. To get all todo.\n4. To delete a todo\n5. To exit\n")
	var userChoise int
	fmt.Scanf("%d", &userChoise)
	return userChoise
}


func main()  {
	var contr controller.Controller
	contr = controller.NewTodoController() 
	// contr := controller2.NewAnotherTypeOfController() 
	choice := 1
	if choice == 1 {
		newTodo, err := contr.Create("test", "test")
		fmt.Println(newTodo, err)
	} else if choice == 2{
		contr.SearchById(1)
	} else {
		contr.SearchByTitle("test")
	}
}
