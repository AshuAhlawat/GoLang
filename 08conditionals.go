package main

import "fmt"


func main(){

	// declaring the age variable
	var age int8
	
	// taking input for value of age
	print("Enter your age : ")
	fmt.Scan(&age)
	print("\n")


	// this type of syntax and indentations are mandatory
	if age<0 {
		print("Wrong Input")
	} else if age<=4 {
		print("You are an Infant")
	} else if age<=12 {
		print("You are a Kid")
	} else if age<18 {
		print("You are a Teen")
	} else if age<60 {
		print("You are an Adult")
	} else {
		print("You are old")
	}

}