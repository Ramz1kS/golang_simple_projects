package io

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"tictactoe/consts"
	"tictactoe/structs"
)

// Gets move coordinates from stdin, returns row and column coordinates
func GetMoveCoordinates(field *structs.Field, stdinReader *bufio.Reader) (int, int) {
	fmt.Printf("Enter coordinates of your move: ")
	for {
		text, _ := stdinReader.ReadString('\n')
		words := strings.Split(strings.TrimSpace(text), " ")
		var row, col int
		if len(words) != 2 {
			fmt.Println("You passed bad coordinates :(")
			continue
		}
		row, err := strconv.Atoi(words[0])
		if err != nil {
			fmt.Println("You passed bad coordinates :(")
			continue
		}
		col, err = strconv.Atoi(words[1])
		if err != nil {
			fmt.Println("You passed bad coordinates :(")
			continue
		}
		if row < 1 || row > 3 || col < 1 || col > 3 {
			fmt.Println("You passed coordinates that are out of range :c")
			continue
		}
		if field[row-1][col-1] == consts.SYM_EMPTY {
			return row, col
		}
	}
}

// Reads string from stdin, trims space and checks if it's empty or not
func ReadStringNonEmpty(stdinReader *bufio.Reader) (string, error) {
	input, err := stdinReader.ReadString('\n')
	if err != nil {
		return input, err
	}
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return input, errors.New("empty input")
	}
	return input, nil
}

// Reads string from stdin, converts it to int automatically
func GetNumberFromInput(stdinReader *bufio.Reader) (int, error) {
	input, err := ReadStringNonEmpty(stdinReader)
	if err != nil {
		return 0, err
	}
	retVal, err := strconv.Atoi(input)
	return retVal, err
}

// Sets new name
func SetNewName(og_name *string, stdinReader *bufio.Reader) error {
	fmt.Printf("Print your new name: ")
	input, err := ReadStringNonEmpty(stdinReader)
	if err != nil {
		return err
	}
	*og_name = input
	fmt.Printf("Now your name is %s\n", *og_name)
	return nil
}


// summons elder gods
func PrintMenu(isCircle bool) {
	fmt.Println("==TIC TAC TOE==")
	fmt.Println("1. Play")
	fmt.Println("2. Set name")
	if isCircle {
		fmt.Println("3. Switch to cross")
	} else {
		fmt.Println("3. Switch to circle")
	}
	fmt.Println("4. Exit")
}
