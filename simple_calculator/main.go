package main

import (
	"bufio"
	"fmt"
	"os"

	"calculator_project/calculator"
)

func main() {
	input := bufio.NewReader(os.Stdin)
	var operationNumber int
	var rc int
	var a float64
	var b float64

	fmt.Println("Это простой калькулятор на Go-шке.")
	fmt.Println("Введите номер операции: ")
	fmt.Println("1. Сложение")
	fmt.Println("2. Вычитание")
	fmt.Println("3. Умножение")
	fmt.Println("4. Деление")

	rc = calculator.GetOperationNumber(input, &operationNumber)
	if rc != calculator.OkCode {
		calculator.PrintErrorMessage(rc)
		os.Exit(rc)
	}

	fmt.Println("Введите первое число: ")
	rc = calculator.GetCalculationNumber(input, &a)
	if rc != calculator.OkCode {
		calculator.PrintErrorMessage(rc)
		os.Exit(rc)
	}
	fmt.Println("Введите второе число: ")
	rc = calculator.GetCalculationNumber(input, &b)
	if rc != calculator.OkCode {
		calculator.PrintErrorMessage(rc)
		os.Exit(rc)
	}

	rc = calculator.CheckSecondNumber(operationNumber, b)
	if rc != calculator.OkCode {
		calculator.PrintErrorMessage(rc)
		os.Exit(rc)
	}

	result := calculator.Calculate(a, b, operationNumber)
	fmt.Printf("Результат вычислений: %f\n", result)
}
