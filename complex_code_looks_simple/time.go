package complexcodelookssimple

import (
	"fmt"
	"time"
)

const (
	inputDate = "2024-05-13 14:30:00"
	timeZone  = "Europe/Moscow"
)

func ParseDate(s string, loc *time.Location) (time.Time, error) {
	date, err := time.ParseInLocation(time.DateTime, s, loc)
	if err != nil {
		return time.Time{}, fmt.Errorf("parse date %q: %w", s, err)
	}

	return date, nil
}

func ExampleParseDate() {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		fmt.Printf("load location %q: %v\n", timeZone, err)
		return
	}

	date, err := ParseDate(inputDate, loc)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Date:", date)
}
