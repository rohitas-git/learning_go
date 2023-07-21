package stuff

import (
	"errors"
	"strconv"
	"time"
)

var Name string = "Derek"

func IntArrToStrArr(intArr []int) []string {
	var strArr []string
	for _, v := range intArr {
		strArr = append(strArr, strconv.Itoa(v))
	}
	return strArr
}

/* ---------------------------------- Date ---------------------------------- */
// Getter/Setter function
// - To hide the data for anyone who tries to use the package
// - For Encapsulation

type Date struct {
	day   int
	month int
	year  int
}

// Setter
func (d *Date) SetDay(day int) error {
	if (day < 1) || (day > 31) {
		return errors.New("incorrect day value")	// Create a new error
	}
	d.day = day
	return nil
}

func (d *Date) SetMonth(month int) error {
	if (month < 1) || (month > 12) {
		return errors.New("incorrect month value")
	}
	d.month = month
	return nil
}

func (d *Date) SetYear(year int) error {
	if (year < 1875) || (year > time.Now().Year()) {
		return errors.New("incorrect year value")
	}
	d.year = year
	return nil
}

// Getter
func (d *Date) Day() int{
	return d.day
}
func (d *Date) Month() int{
	return d.month
}
func (d *Date) Year() int{
	return d.year
}