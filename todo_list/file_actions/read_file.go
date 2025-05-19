package file_actions

import (
	"strings"
	"errors"
	"os"
	"bufio"
	"todo_list/structs"
	"todo_list/misc"
)

func ShouldReadFile(text string) bool {
	text = strings.TrimSpace(text)
	return text != ""
}

func ReadData(filename *string, data structs.TasksList) error {
	if *filename == "" {
		return errors.New("empty filename")
	}
	file, err := os.Open(*filename)
	if err != nil {
		return err
	}
	fileScanner := bufio.NewScanner(file)
	hasDate := false
	var currDate structs.Date
	for fileScanner.Scan() {
		text := fileScanner.Text()
		text = strings.Trim(text, " \t\r\n")
		if text == "" {
			continue
		}
		if text[0:5] == "date " {
			hasDate = true
			currDate, err = misc.StringToDate(text[5:])
			if err != nil {
				return err
			}
			_, exists := data[currDate]
			if !exists {
				data[currDate] = []string{}
			}
		} else if hasDate {
			data[currDate] = append(data[currDate], text)
		} else {
			return errors.New("no date specified for a task list")
		}
	}
	return nil
}