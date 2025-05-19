package list_actions

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"todo_list/structs"
)

func SaveTable(data structs.TasksList, stdinReader *bufio.Reader) error {
	fmt.Printf("Please enter a name of a file: ")
	text, _ := stdinReader.ReadString('\n')
	text = strings.TrimSpace(text)
	if len(text) == 0 {
		return errors.New("empty file name")
	}
	file, err := os.Create(text)
  if err != nil {
    return err
  }
  for date, tasks := range data {
		fmt.Fprintf(file, "date %d.%d.%d\n", date.Day, date.Month, date.Year)
		for _, task := range tasks {
			fmt.Fprintf(file, "%s \n", task)
		}
	}
	file.Close()
	return nil
}