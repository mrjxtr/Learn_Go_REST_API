package main

import (
	"fmt"

	"github.com/mrjxtr/Learn_Go_REST_API/learn/section4/practice/note"
)

func main() {
	title, content := note.GetNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Printf("Error creating note: %v\n", err)
		return
	}

	userNote.OutputNote()

	err = userNote.Save()
	if err != nil {
		fmt.Printf("Saving note failed: %v", err)
		return
	}

	fmt.Println("Note Saved!")
}
