package io

import (
	"tictactoe/structs"
	"fmt"
)

// Prints field
func PrintField(field *structs.Field) {
	fmt.Println("====CURRENT PLAYING FIELD====")
	for _, row := range field {
		for _, elem := range row {
			fmt.Printf("%c ", elem)
		}
		fmt.Printf("\n")
	}
}