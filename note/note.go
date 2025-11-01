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
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Printf("Your note titled==> %v has the following content:\n%v\n", note.Title, note.Content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")     // create file name replace space to  _
	fileName = "note/" + strings.ToLower(fileName) + ".json" // save in note folder

	data, err := json.Marshal(note) //export text to Json format

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input.")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
