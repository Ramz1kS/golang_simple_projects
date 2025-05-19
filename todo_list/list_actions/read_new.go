package list_actions

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
	"todo_list/file_actions"
	"todo_list/structs"
)

func ReadNewFile(data *(structs.TasksList), stdinReader *bufio.Reader) error {
	fmt.Printf("Please enter a name of a file: ")
	text, _ := stdinReader.ReadString('\n')
	text = strings.TrimSpace(text)
	if len(text) == 0 {
		return errors.New("empty file name")
	}
	data_new := structs.TasksList{}
	err := file_actions.ReadData(&text, data_new)
	if err == nil {
		*data = data_new
	}
	return nil
}