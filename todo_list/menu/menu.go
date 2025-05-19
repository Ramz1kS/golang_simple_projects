package menu

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"todo_list/codes"
)

func checkIfChoiceValid(choice int) bool {
	return choice > 0 && choice <= codes.CodeExit
}

func GetChoice(stdinReader *bufio.Reader) int {
	fmt.Println("======MENU======")
	fmt.Println("1. Print all the tasks")
	fmt.Println("2. Add new task")
	fmt.Println("3. Mark task as completed")
	fmt.Println("4. Open new file")
	fmt.Println("5. Save file")
	fmt.Println("6. Exit")
	fmt.Printf("Enter an option number: ")
	for {
		text, _ := stdinReader.ReadString('\n')
		text = strings.TrimSpace(text)
		choice, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Please input a number and nothing more")
		} else if !checkIfChoiceValid(choice) {
			fmt.Println("Choice number is out of range")
		} else {
			return choice
		}
	}
}