package structs

type Date struct {
	Day int;
	Month int;
	Year int
}

type TasksList map[Date]([]string)