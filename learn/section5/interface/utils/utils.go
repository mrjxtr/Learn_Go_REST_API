package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Saver interface {
	Save() error
}

type Outputer interface {
	Output()
}

type Output interface {
	Saver
	Outputer
}

// Get note data from inputs
func GetNoteData() (string, string) {
	title := getUserInput("\nNote Title: ")
	content := getUserInput("Note Content: \n\n")

	return title, content
}

// Get todo data from inputs
func GetTodoData() string {
	return getUserInput("\nTodo: ")
}

// Get multi-word user input from terminal
func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
}

// A generic way of saving notes or todos using the Saver interface
func SaveData(data Saver) error {
	err := data.Save()
	if err != nil {
		return errors.New("failed to save data")
	} else {
		fmt.Println("Data Saved!")
	}

	return nil
}

func OutputData(data Output) error {
	data.Output()
	return SaveData(data)
}
