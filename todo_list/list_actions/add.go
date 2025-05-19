package list_actions

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
	"todo_list/misc"
	"todo_list/structs"
)

func AddEntry(list structs.TasksList, stdinReader *bufio.Reader) (error) {
	var date structs.Date
	fmt.Printf("Please enter a date in DD.MM.YYYY format: ")
	text, _ := stdinReader.ReadString('\n')
	text = strings.TrimSpace(text)
	var err error
	date, err = misc.StringToDate(text)
	if err != nil {
		return err
	}
	fmt.Printf("Please enter a new task: ")
	text, _ = stdinReader.ReadString('\n')
	text = strings.TrimSpace(text)
	if len(text) == 0 {
		return errors.New("empty input")
	}
	list[date] = append(list[date], text)
	return nil
}