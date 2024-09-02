package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"example.com/playground/lib"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (*Note, error) {
	if title == "" {
		return nil, errors.New("Title is required")
	} else if content == "" {
		return nil, errors.New("Content is required")
	}

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Note has title: %v\n content: %v\n created at: %v",
		note.Title,
		note.Content,
		note.CreatedAt.Format(time.DateTime))
}

func (note Note) Save() error {

	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName)

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName+".json", json, 0644)
}

func CreateNote() *Note {

	var title, content string

	lib.GetUserInput("Enter a note title:", &title)
	lib.GetUserInput("Enter note content", &content)

	note, err := New(title, content)

	if err != nil {
		panic(fmt.Sprintf("%v\n", err))
	}

	return note
}
