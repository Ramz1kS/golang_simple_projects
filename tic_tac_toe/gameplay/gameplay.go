package gameplay

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"tictactoe/consts"
	t_io "tictactoe/io"
	"tictactoe/structs"
)

// Checks if current round is over or not
func CheckEnd(field *structs.Field) structs.WinInfo {
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
func MarkStep(field *structs.Field, symbol rune, row int, col int) error {
	if field[row-1][col-1] != consts.SYM_EMPTY {
		return errors.New("can't mark step on a place that was marked already")
	}
	field[row-1][col-1] = symbol
	return nil
}

// Initialises playing field with "empty" symbols
func InitField() structs.Field {
	return structs.Field{
		{consts.SYM_EMPTY, consts.SYM_EMPTY, consts.SYM_EMPTY},
		{consts.SYM_EMPTY, consts.SYM_EMPTY, consts.SYM_EMPTY},
		{consts.SYM_EMPTY, consts.SYM_EMPTY, consts.SYM_EMPTY},
	}
}

func MarkWinArea(field *structs.Field, info *structs.WinInfo) {
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
		t_io.PrintField(&field)
		moveRow, moveCol := t_io.GetMoveCoordinates(&field, stdinReader)
		err := MarkStep(&field, currSymbol, moveRow, moveCol)
		if err != nil {
			log.Fatal(err)
		}
		winInfo = CheckEnd(&field)
	}
	if CheckEnd(&field).DidWin {
		MarkWinArea(&field, &winInfo)
		t_io.PrintField(&field)
		fmt.Println("YOU WON! CONGRATS!")
	} else {
		t_io.PrintField(&field)
		fmt.Println("DRAW!")
	}
}
