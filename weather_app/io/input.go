package w_io

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"weather/colors"
)

func GetChoice(welcome string, max int, stdinReader *bufio.Reader) (int, error) {
	var result int
	if max < 1 {
		return result, errors.New("max value is less than 1")
	}
	fmt.Printf("%s", welcome)
	text, err := stdinReader.ReadString('\n')
	if err == nil {
		text = strings.TrimSpace(text)
		result, err = strconv.Atoi(text)
		if err == nil && !(1 <= result && result <= max) {
			err = errors.New("choice is out of range")
		}
	}
	return result, err
}

func GetMeasurementMethod(stdinReader *bufio.Reader) (string, error) {
	methods := [...]string{"standard", "metric", "imperial"}
	fmt.Println("MEASUREMENT METHODS")
	fmt.Printf("1. %sStandard%s\n", colors.Cyan, colors.Reset)
	fmt.Printf("2. %sMetric%s\n", colors.Cyan, colors.Reset)
	fmt.Printf("3. %sImperial%s\n", colors.Cyan, colors.Reset)
	choice, err := GetChoice("Please select a measurement method: ", len(methods), stdinReader)
	if err == nil {
		return methods[choice - 1], err
	} else {
		return "", err
	}
}