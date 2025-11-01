package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/first/note"
)

func NoteBook() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note succeeded!")
}

func getNoteData() (string, string) {
	title := getInput("Note title:")
	content := getInput("Note content:")

	return title, content
}

func getInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin) //look like scan but can handle text and /n

	text, err := reader.ReadString('\n') //it means end include with /n

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n") //remove /n from text
	text = strings.TrimSuffix(text, "\r") // remove /r from text

	return text
}
