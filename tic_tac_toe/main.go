package main

import (
	"bufio"
	"fmt"
	"os"
	t_io "tictactoe/io"
	"tictactoe/consts"
	"tictactoe/gameplay"
)

func main() {
	name := "Unknown Slayer"
	currSymbol := consts.SYM_CROSS
	stdinReader := bufio.NewReader(os.Stdin)
	choice := consts.ACTION_PLAY
	for choice != consts.ACTION_EXIT {
		t_io.PrintMenu(currSymbol == consts.SYM_ZERO)
		fmt.Printf("Choose an option: ")
		choice, err := t_io.GetNumberFromInput(stdinReader)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Try again")
			continue
		}

		switch choice {
			case consts.ACTION_PLAY:
				gameplay.StartSinglePlayer(currSymbol, stdinReader)
			case consts.ACTION_RENAME:
				err = t_io.SetNewName(&name, stdinReader)
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
