package main

import (
	"fmt"

	"github.com/mrjxtr/Learn_Go_REST_API/learn/section5/interface/note"
	"github.com/mrjxtr/Learn_Go_REST_API/learn/section5/interface/todo"
	"github.com/mrjxtr/Learn_Go_REST_API/learn/section5/interface/utils"
)

func main() {
	title, content := utils.GetNoteData()
	todoText := utils.GetTodoData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Printf("Error creating note: %v\n", err)
		return
	}

	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Printf("Error creating todo: %v\n", err)
		return
	}

	// i: Best way
	err = utils.OutputData(userNote)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	// i: Better way
	// userNote.Output()
	// err = utils.SaveData(userNote) // Using generic save method
	// if err != nil {
	// 	fmt.Printf("Error: %s", err)
	// 	return
	// }
	// i: Long way
	// err = userNote.Save()
	// if err != nil {
	// 	fmt.Printf("Saving note failed: %v", err)
	// 	return
	// } else {
	// 	fmt.Println("Note Saved!")
	// }

	// i: Best way
	err = utils.OutputData(userTodo)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	// i: Better way
	// userTodo.Output()
	// err = utils.SaveData(userTodo) // Using generic save method
	// if err != nil {
	// 	fmt.Printf("Error: %s", err)
	// 	return
	// }
	// i: Long way
	// err = userTodo.Save()
	// if err != nil {
	// 	fmt.Printf("Saving todo failed: %v", err)
	// 	return
	// } else {
	// 	fmt.Println("Todo Saved!")
	// }
}
