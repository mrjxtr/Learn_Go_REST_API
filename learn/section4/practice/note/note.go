package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

// Note struct constructor
func New(title, content string) (*Note, error) {
	if title == "" {
		return nil, errors.New("title should not be empty")
	}
	if content == "" {
		return nil, errors.New("content should not be empty")
	}

	return &Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}

// Method to output note title and content
func (n *Note) OutputNote() {
	fmt.Printf(
		"\n===========================\nThis is your note...\nTitle: %s\n\n%s\n\n",
		n.title,
		n.content,
	)
	fmt.Printf("Time Created: %v\n", n.createdAt.Format("2006/02/01"))
}

func (n *Note) Save() error {
	fileDir := "./learn/section4/practice/data/"

	fileName := strings.ReplaceAll(n.title, " ", "_") // repace all spaces with "_"
	fileName = strings.ToLower(fileName) + ".json"    // all characters to lower case

	// Create a temporary struct to format CreatedAt
	// i: This also allows me to avoid making the Note struct fiels public
	temp := struct {
		Title     string `json:"title"`
		Content   string `json:"content"`
		CreatedAt string `json:"created_at"`
	}{
		Title:     n.title,
		Content:   n.content,
		CreatedAt: n.createdAt.Format("2006/02/01"), // format date for the json file
	}

	data, err := json.MarshalIndent(temp, "", "  ") // pretty print with 2 spaces indent
	if err != nil {
		return err
	}

	return os.WriteFile(fileDir+fileName, data, 0644)
}
