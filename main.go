package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {

	var date string
	fmt.Printf("Enter a date: ")
	fmt.Scanln(&date)

	str := strings.Split(date, "/")

	day, _ := strconv.Atoi(str[0])
	month, _ := strconv.Atoi(str[1])
	year, _ := strconv.Atoi(str[2])

	/*myMap := make(map[string][]int)

	myMap["Sunday"] = []int{}
	myMap["Monday"] = []int{}
	myMap["Tuesday"] = []int{}
	myMap["Wednesday"] = []int{}
	myMap["Thursday"] = []int{}
	myMap["Friday"] = []int{}
	myMap["Saturday"] = []int{} */

	x := time.January

	switch month {
	case 1:
		x = time.January
	case 2:
		x = time.February
	case 3:
		x = time.March
	case 4:
		x = time.April
	case 5:
		x = time.May
	case 6:
		x = time.June
	case 7:
		x = time.July
	case 8:
		x = time.August
	case 9:
		x = time.September
	case 10:
		x = time.October
	case 11:
		x = time.November
	case 12:
		x = time.December
	}

	firstDayOfMonth := time.Date(year, x, 1, 1, 1, 1, 1, time.UTC)
	firstDayOfMonthWeekday := firstDayOfMonth.Weekday().String()
	monthOfFirstDayOfMonth := firstDayOfMonth.Month().String()

	myMap := make(map[string]int)

	myMap["January"] = 31
	if ((year%4 == 0) && (year%100 != 0)) || (year%400 == 0) {
		myMap["February"] = 29
	} else {
		myMap["February"] = 28
	}
	myMap["March"] = 31
	myMap["April"] = 30
	myMap["May"] = 31
	myMap["June"] = 30
	myMap["July"] = 31
	myMap["August"] = 31
	myMap["September"] = 30
	myMap["October"] = 31
	myMap["November"] = 30
	myMap["December"] = 31

	/*newDate := myDate
	newMonth := month
	newYear := year

	for newMonth <= month && newYear <= year {
		myMap[newDate.Weekday().String()] = append(myMap[newDate.Weekday().String()], newDate.Day())
		newDate = newDate.Add(24 * time.Hour)
		newMonth = int(newDate.Month())
		newYear = newDate.Year()
	}

	for key, value := range myMap {
		fmt.Println(key, ": ", value)
	}*/

	newMap := make(map[string]int)

	newMap["Sunday"] = 0
	newMap["Monday"] = 1
	newMap["Tuesday"] = 2
	newMap["Wednesday"] = 3
	newMap["Thursday"] = 4
	newMap["Friday"] = 5
	newMap["Saturday"] = 6

	arr := [6][7]int{}

	lim := myMap[monthOfFirstDayOfMonth]
	counter := 1

	for j := newMap[firstDayOfMonthWeekday]; j < 7; j += 1 {
		arr[0][j] = counter
		counter += 1
	}

	for i := 1; i < 6 && counter <= lim; i += 1 {
		for j := 0; j < 7 && counter <= lim; j += 1 {
			arr[i][j] = counter
			counter += 1
		}
	}

	for i := 0; i < 6; i += 1 {
		for j := 0; j < 7; j += 1 {
			if arr[i][j] < day {
				arr[i][j] = 0
			}
		}
	}

	fmt.Println("\nPlease find the required Calender below\n")

	fmt.Println("Sun    Mon     Tues    Wed     Thurs    Fri    Sat")
	for i := 0; i < 6; i += 1 {
		for j := 0; j < 7; j += 1 {
			if arr[i][j] < 10 {
				fmt.Printf("%d       ", arr[i][j])
			} else {
				fmt.Printf("%d      ", arr[i][j])
			}
		}
		fmt.Println()
	}

}
