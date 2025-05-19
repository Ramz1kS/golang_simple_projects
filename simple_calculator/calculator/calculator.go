package calculator

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Calculate(a float64, b float64, operNum int) float64 {
	switch operNum {
	case AdditionCode:
		return a + b
	case SubtractionCode:
		return a - b
	case MultiplicationCode:
		return a * b
	case DivisionCode:
		return a / b
	}
	return a + b
}

func isFloatEqual(a float64, b float64) bool {
	return math.Abs(a - b) < 0.000001
}

func PrintErrorMessage(rc int) {
	switch rc {
	case BadInputCode:
		fmt.Println("Некорректный ввод :(")
	case ZeroDivisionCode:
		fmt.Println("Деление на ноль невозможно")
	default:
		fmt.Printf("Ошибка с кодом: %d\n", rc)
	}
}

func GetCalculationNumber(input *bufio.Reader, number *float64) int {
	text, err := input.ReadString('\n')
	text = strings.TrimSpace(text)
	if err != nil {
		return BadInputCode
	}
	numberValue, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return BadInputCode
	}
	*number = numberValue
	return OkCode
}

func GetOperationNumber(input *bufio.Reader, number *int) int {
	text, err := input.ReadString('\n')
	text = strings.TrimSpace(text)
	if err != nil {
		return BadInputCode
	}
	numberValue, err := strconv.Atoi(text)
	if err != nil {
		return BadInputCode
	}
	*number = numberValue
	return OkCode
}

func CheckSecondNumber(operationNumber int, number float64) int {
	if operationNumber == DivisionCode && isFloatEqual(number, 0) {
		return ZeroDivisionCode
	}
	return OkCode
}