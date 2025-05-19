package misc

import (
	"errors"
	"sort"
	"strconv"
	"strings"
	"todo_list/structs"
)

func CheckIfNumInRange(num, a, b int) (bool, error) {
	if a > b {
		return false, errors.New("while checking if a <= num <= b, a cannot be greater than b")
	}
	return num >= a && num <= b, nil
}

func checkIfDayInRange(day int, month int) (bool, error) {
	if month <= 0 || month >= 13 {
		return false, errors.New("month out of range")
	}
	dayTable := map[int]int{
		1:  31,
		2:  28, // без учета високосного года
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
	return CheckIfNumInRange(day, 1, dayTable[month])
}

func StringToDate(text string) (structs.Date, error) {
	var retDate structs.Date
	datePartsStrings := strings.Split(text, ".")
	if len(datePartsStrings) != 3 {
		return retDate, errors.New("bad date length")
	}
	day_val, err := strconv.Atoi(datePartsStrings[0])
	if err != nil {
		return retDate, errors.New("bad day in date")
	}
	month_val, err := strconv.Atoi(datePartsStrings[1])
	isMonthCorrect, _ := CheckIfNumInRange(month_val, 1, 12)
	if err != nil || !isMonthCorrect {
		return retDate, errors.New("bad month in date")
	}
	isDayCorrect, _ := checkIfDayInRange(day_val, month_val)
	if !isDayCorrect {
		return retDate, errors.New("bad day in date")
	}
	year_val, err := strconv.Atoi(datePartsStrings[2])
	if err != nil || year_val < 1 {
		return retDate, errors.New("bad year in date")
	}
	retDate = structs.Date{Day: day_val, Month: month_val, Year: year_val}
	return retDate, nil
}

func GetSortedDates(list structs.TasksList) []structs.Date {
	dates := []structs.Date{}
	for key := range list {
		if len(list[key]) != 0 {
			dates = append(dates, key)
		}
	}
	sort.Slice(dates, func(i, j int) bool {
		if dates[i].Year != dates[j].Year {
			return dates[i].Year < dates[j].Year
		}
		if dates[i].Month != dates[j].Month {
			return dates[i].Month < dates[j].Month
		}
		return dates[i].Day < dates[j].Day
	})
	return dates
}
