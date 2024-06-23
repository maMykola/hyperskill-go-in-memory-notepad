package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Notepad struct {
	Capacity int
	Notes    []string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	notepad := NewNotepad(getNumberOfNotes())

	for {
		fmt.Print("Enter a command and data: ")

		scanner.Scan()
		command, data := splitInput(scanner.Text())

		switch command {
		case "exit":
			fmt.Println("[Info] Bye!")
			os.Exit(0)
		case "create":
			notepad.AddNote(data)
		case "list":
			notepad.PrintNotes()
		case "clear":
			notepad.ClearNotes()
		case "update":
			pos, text := splitInput(data)
			notepad.UpdateNote(pos, text)
		case "delete":
			pos, _ := splitInput(data)
			notepad.DeleteNode(pos)
		default:
			fmt.Println("[Error] Unknown command")
		}

		fmt.Println()
	}
}

func splitInput(input string) (string, string) {
	data := strings.SplitN(input, " ", 2)

	if len(data) == 2 {
		return data[0], strings.TrimSpace(data[1])
	}

	return data[0], ""
}

func getNumberOfNotes() int {
	var num int

	fmt.Print("Enter the maximum number of notes: ")
	fmt.Scan(&num)
	fmt.Println()

	return num
}

func NewNotepad(size int) *Notepad {
	notepad := new(Notepad)
	notepad.Capacity = size
	notepad.Notes = make([]string, 0, size)
	return notepad
}

func (np *Notepad) AddNote(note string) {
	if len(note) == 0 {
		fmt.Println("[Error] Missing note argument")
	} else if len(np.Notes) == np.Capacity {
		fmt.Println("[Error] Notepad is full")
	} else {
		fmt.Println("[OK] The note was successfully created")
		np.Notes = append(np.Notes, note)
	}
}

func (np *Notepad) UpdateNote(line, note string) {
	pos, err := np.GetPosition(line)

	switch true {
	case err != nil:
		fmt.Println(err)
	case note == "":
		fmt.Println("[Error] Missing note argument")
	case pos > np.Capacity:
		fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", pos+1, np.Capacity)
	case len(np.Notes) <= pos:
		fmt.Println("[Error] There is nothing to updated")
	default:
		np.Notes[pos] = note
		fmt.Printf("[OK] The note at position %d was successfully updated\n", pos+1)
	}
}

func (np *Notepad) DeleteNode(line string) {
	pos, err := np.GetPosition(line)

	switch true {
	case err != nil:
		fmt.Println(err)
	case pos > np.Capacity:
		fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", pos+1, np.Capacity)
	case len(np.Notes) <= pos:
		fmt.Println("[Error] There is nothing to deleted")
	default:
		copy(np.Notes[pos:], np.Notes[pos+1:])
		np.Notes = np.Notes[:len(np.Notes)-1]
		fmt.Printf("[OK] The note at position %d was successfully deleted\n", pos+1)
	}
}

func (np *Notepad) PrintNotes() {
	if len(np.Notes) == 0 {
		fmt.Println("[Info] Notepad is empty")
		return
	}

	for i, note := range np.Notes {
		fmt.Printf("[Info] %d: %s\n", i+1, note)
	}
}

func (np *Notepad) ClearNotes() {
	np.Notes = np.Notes[:0]
	fmt.Println("[OK] All notes were successfully deleted")
}

func (np *Notepad) GetPosition(line string) (int, error) {
	if line == "" {
		return 0, errors.New("[Error] Missing position argument")
	}

	pos, err := strconv.Atoi(line)
	if err != nil {
		return 0, fmt.Errorf("[Error] Invalid position: %s", line)
	}

	return pos - 1, nil
}
