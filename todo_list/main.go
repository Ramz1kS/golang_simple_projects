package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"todo_list/codes"
	"todo_list/file_actions"
	"todo_list/menu"
	"todo_list/structs"
	"todo_list/list_actions"
)

func main() {
	data := make(structs.TasksList)
	var err error
	stdinReader := bufio.NewReader(os.Stdin)

	fmt.Println("TODO-list!")
	fmt.Println("Please write a file name with your tasks. If you don't have a file, just please Enter")
	filename, _ := stdinReader.ReadString('\n')
	filename = strings.TrimSpace(filename)
	if file_actions.ShouldReadFile(filename) {
		err = file_actions.ReadData(&filename, data)
		if err != nil {
			log.Fatal(err)
		}
	}

	var choice int = codes.CodePrint
	for choice != codes.CodeExit {
		if err != nil {
			fmt.Println(err)
		}
		err = nil
		choice = menu.GetChoice(stdinReader)
		switch choice {
		case codes.CodePrint:
			list_actions.PrintList(data)
		case codes.CodeAdd:
			err = list_actions.AddEntry(data, stdinReader)
		case codes.CodeRemove:
			err = list_actions.RemoveTask(data, stdinReader)
		case codes.CodeOpen:
			err = list_actions.ReadNewFile(&data, stdinReader)
		case codes.CodeSave:
			err = list_actions.SaveTable(data, stdinReader)
		case codes.CodeExit:
			return
		default:
			fmt.Println("Not done yet!")
		}
	} 
}
