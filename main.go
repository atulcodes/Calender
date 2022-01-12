package main

import "fmt"

func main() {
	var day, month, year int
	fmt.Println("Enter the Day: ")
	fmt.Scanln(&day)
	fmt.Println("Enter the Month: ")
	fmt.Scanln(&month)
	fmt.Println("Enter the year: ")
	fmt.Scanln(&year)
	fmt.Println("Entered Date: ", day, "/", month, "/", year)
}
