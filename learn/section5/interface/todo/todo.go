package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	text      string
	createdAt time.Time
}

// Todo struct constructor
func New(text string) (*Todo, error) {
	if text == "" {
		return nil, errors.New("todo should not be empty")
	}

	return &Todo{
		text:      text,
		createdAt: time.Now(),
	}, nil
}

// Method to output todo title and content
func (t *Todo) Output() {
	fmt.Printf(
		"\n===========================\nThis is your todo...\n%s\n\n",
		t.text,
	)
	fmt.Printf("Time Created: %v\n", t.createdAt.Format("2006/02/01"))
}

// Method to save todo
func (t *Todo) Save() error {
	fileDir := "./learn/section5/interface/data/"

	fileName := "todo.json"

	// Create a temporary struct to format CreatedAt
	// i: This also allows me to avoid making the todo struct fiels public
	temp := struct {
		Text      string `json:"text"`
		CreatedAt string `json:"created_at"`
	}{
		Text:      t.text,
		CreatedAt: t.createdAt.Format("2006/02/01"), // format date for the json file
	}

	data, err := json.MarshalIndent(temp, "", "  ") // pretty print with 2 spaces indent
	if err != nil {
		return err
	}

	return os.WriteFile(fileDir+fileName, data, 0644)
}
