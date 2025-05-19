package list_actions

import (
	"fmt"
	"todo_list/structs"
	"todo_list/misc"
)

func PrintList(list structs.TasksList) {
	if len(list) == 0 {
		fmt.Println("You don't have anything in your todo-list")
		return
	}
	dates := misc.GetSortedDates(list)
	for _, key := range dates {
		fmt.Printf("DATE: %d.%d.%d\n", key.Day, key.Month, key.Year)
		for _, task := range list[key] {
			fmt.Printf("- %s\n", task)
		}
	}
}
