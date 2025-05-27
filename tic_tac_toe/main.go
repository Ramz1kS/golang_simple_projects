package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"tictactoe/consts"
	"tictactoe/structs"
)

type Field [3][3]rune

// Prints field
func PrintField(field *Field) {
	fmt.Println("====CURRENT PLAYING FIELD====")
	for _, row := range field {
		for _, elem := range row {
			fmt.Printf("%c ", elem)
		}
		fmt.Printf("\n")
	}
}

// Gets move coordinates from stdin, returns row and column coordinates
func GetMoveCoordinates(field *Field, stdinReader *bufio.Reader) (int, int) {
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

// Checks if current round is over or not
func CheckEnd(field *Field) structs.WinInfo {
	// horizontal check
	for i, row := range field {
		if row[0] != consts.SYM_EMPTY && row[0] == row[1] && row[1] == row[2] {
			return structs.WinInfo{DidWin: true, DidEnd: true, Symbol: row[0], Direction: consts.DIR_HORIZ, StartRow: i}
		}
	}
	// vertical check
	for i := range len(field[0]) {
		if field[0][i] != consts.SYM_EMPTY && field[0][i] == field[1][i] && field[1][i] == field[2][i] {
			return structs.WinInfo{DidWin: true, DidEnd: true, Symbol: field[0][i], Direction: consts.DIR_VERT, StartCol: i}
		}
	}
	// angled check
	if field[0][0] != consts.SYM_EMPTY && field[0][0] == field[1][1] && field[1][1] == field[2][2] {
		return structs.WinInfo{DidWin: true, DidEnd: true, Symbol: field[0][0], Direction: consts.DIR_ANGL_DWN}
	}
	if field[2][0] != consts.SYM_EMPTY && field[2][0] == field[1][1] && field[1][1] == field[0][2] {
		return structs.WinInfo{DidWin: true, DidEnd: true, Symbol: field[0][0], Direction: consts.DIR_ANGL_UP}
	}
	// check if there's no empty space
	for _, row := range field {
		for _, col := range row {
			if col == consts.SYM_EMPTY {
				return structs.WinInfo{DidWin: false, DidEnd: false}
			}
		}
	}
	return structs.WinInfo{DidWin: false, DidEnd: true}
}

// Function to mark step
func MarkStep(field *Field, symbol rune, row int, col int) error {
	if field[row-1][col-1] != consts.SYM_EMPTY {
		return errors.New("can't mark step on a place that was marked already")
	}
	field[row-1][col-1] = symbol
	return nil
}

// Initialises playing field with "empty" symbols
func InitField() Field {
	return Field{
		{consts.SYM_EMPTY, consts.SYM_EMPTY, consts.SYM_EMPTY},
		{consts.SYM_EMPTY, consts.SYM_EMPTY, consts.SYM_EMPTY},
		{consts.SYM_EMPTY, consts.SYM_EMPTY, consts.SYM_EMPTY},
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

func MarkWinArea(field *Field, info *structs.WinInfo) {
	if !info.DidEnd || !info.DidWin {
		fmt.Println("Nothing to mark here lil bro")
		return
	}
	// consume the milk chalice
	if info.Direction == consts.DIR_HORIZ {
		for i := range 3 {
			field[info.StartRow][i] = '-'
		}
	} else if info.Direction == consts.DIR_VERT {
		for i := range 3 {
			field[i][info.StartCol] = '|'
		}
	} else if info.Direction == consts.DIR_ANGL_DWN {
		for i := range 3 {
			field[i][i] = '\\'
		}
	} else {
		for i := range 3 {
			field[2 - i][i] = '/'
		}
	}
}

func StartSinglePlayer(currSymbol rune, stdinReader *bufio.Reader) {
	field := InitField()
	winInfo := CheckEnd(&field)
	for !winInfo.DidEnd {
		PrintField(&field)
		moveRow, moveCol := GetMoveCoordinates(&field, stdinReader)
		err := MarkStep(&field, currSymbol, moveRow, moveCol)
		if err != nil {
			log.Fatal(err)
		}
		winInfo = CheckEnd(&field)
	}
	if CheckEnd(&field).DidWin {
		MarkWinArea(&field, &winInfo)
		PrintField(&field)
		fmt.Println("YOU WON! CONGRATS!")
	} else {
		PrintField(&field)
		fmt.Println("DRAW!")
	}
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

func main() {
	name := "Unknown Slayer"
	currSymbol := consts.SYM_CROSS
	stdinReader := bufio.NewReader(os.Stdin)
	choice := consts.ACTION_PLAY
	for choice != consts.ACTION_EXIT {
		PrintMenu(currSymbol == consts.SYM_CROSS)
		fmt.Printf("Choose an option: ")
		choice, err := GetNumberFromInput(stdinReader)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Try again")
			continue
		}

		switch choice {
			case consts.ACTION_PLAY:
				StartSinglePlayer(currSymbol, stdinReader)
			case consts.ACTION_RENAME:
				err = SetNewName(&name, stdinReader)
				if err != nil {
					fmt.Println(err)
				}
			case consts.ACTION_SWITCH_SYM:
				if currSymbol == consts.SYM_CROSS {
					currSymbol = consts.SYM_ZERO
				} else {
					currSymbol = consts.SYM_CROSS
				}
				fmt.Println("Switched successfully")
			case consts.ACTION_EXIT:
				fmt.Println("Thanks for playing! :D")
				os.Exit(0)
			default:
				fmt.Println("Out of range... Try again")
		}
	}
}
