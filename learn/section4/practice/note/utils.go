package note

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Get note data from inputs
func GetNoteData() (string, string) {
	title := getUserInput("Note Title: ")
	content := getUserInput("Note Content: \n\n")

	return title, content
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
