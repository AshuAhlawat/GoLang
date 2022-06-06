package main

import "fmt"

func main() {

	var num int8 // initialized to 0

	print("Week Day number :")
	fmt.Scan(&num)

	var day string // initialized to  ""

	switch num{ // starts from curly brackets
		case 1:
			day = "Monday"
		case 2:
			day = "Tuesday"
		case 3:
			day = "Wednesday"
		case 4:
			day = "Thursday"
		case 5:
			day = "Friday"
		case 6:
			day = "Saturday"
		case 0,7: // can give multiple cases so if num == 0 or 7
			day = "Sunday"
		default: // default case to run if no case is true
			print("\nThats not a day")
	
	}

	if day != "" {
		print("\nThe day name is ",day)
	}


}