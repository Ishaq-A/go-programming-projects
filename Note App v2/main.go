package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

// Embedded Interface
type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo Text: ")
	
	todo, err := todo.New(todoText)
	
	if err != nil {
		fmt.Println(err)
		return
	}
	
	userNote, err := note.New(title, content)
	
	if err != nil {
		fmt.Println(err)
		return
	}
	
	err = outputData(todo)
	
	if err != nil {
		fmt.Println(err)
		return
	}
	
	outputData(userNote)
	// Error handling not required as program ends here.
	printSomething(100000)
}

// Example to Show "any" Type
func printSomething(value interface{}) {
	typedVal, ok := value.(int)
	
	if ok {
		fmt.Println("Integer:", typedVal)
	}
}

func outputData(data outputtable) error {
	data.Display()
	saveData(data)
	
	return nil
}

func saveData(data saver) error {
	err := data.Save()
	
	if err != nil {
		fmt.Println("saving note failed.")
		return err
	}
	
	fmt.Println("Saving the note succeeded")
	
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")
	
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	
	reader := bufio.NewReader(os.Stdin)
	
	text, err := reader.ReadString('\n')
	
	if err != nil {
		return ""
	}
	
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	
	return text
}


