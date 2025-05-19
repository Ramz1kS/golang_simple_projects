package list_actions

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"todo_list/misc"
	"todo_list/structs"
)

func getDateNumber(dates []structs.Date, stdinReader *bufio.Reader) (int, error) {
	fmt.Printf("Enter the date number: ")
	text, _ := stdinReader.ReadString('\n')
	text = strings.TrimSpace(text)
	choice, err := strconv.Atoi(text)
	if err != nil {
		return 0, err
	} 
	isNumberGood, err := misc.CheckIfNumInRange(choice, 1, len(dates))
	if err != nil {
		return 0, err
	}
	if !isNumberGood {
		return 0, errors.New("out of range choice")
	}
	return choice, nil
}

func deleteTask(tasks *structs.TasksList, date structs.Date, stdinReader *bufio.Reader) (error) {
	fmt.Printf("Enter the task number: ")
	text, _ := stdinReader.ReadString('\n')
	text = strings.TrimSpace(text)
	choice, err := strconv.Atoi(text)
	if err != nil {
		return err
	} 
	isNumberGood, err := misc.CheckIfNumInRange(choice, 1, len((*tasks)[date]))
	if err != nil {
		return err
	} 
	if !isNumberGood {
		return errors.New("out of range choice")
	}
	tasksForDate := (*tasks)[date]
	(*tasks)[date] = append(tasksForDate[:choice-1], tasksForDate[choice:]...)
  if len((*tasks)[date]) == 0 {
    delete(*tasks, date)
  }
	return nil
}

func RemoveTask(list structs.TasksList, stdinReader *bufio.Reader) error {
	dates := misc.GetSortedDates(list)
	if len(dates) == 0 {
		return errors.New("no tasks in current todo-list")
	}
	fmt.Println("Available dates: ")
	for index, val := range dates {
		fmt.Printf("%d. %d.%d.%d\n", index + 1, val.Day, val.Month, val.Year)
	}
	choice, err := getDateNumber(dates, stdinReader)
	if err != nil {
		return nil
	}
	tasksForDate := list[dates[choice - 1]]
	for index, val := range tasksForDate {
		fmt.Printf("%d. %s\n", index + 1, val)
	}
	err = deleteTask(&list, dates[choice - 1], stdinReader)
	return err
}