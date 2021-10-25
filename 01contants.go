package main

import "fmt"

// must be declared and assigned in same line
// can be out of function
const PI = 3.14

func main() {
	// convention to have all capital chars
	const NAME string = "Constant String"

	const (
		A      = 3.4
		B bool = true
		C      = 21
	)

	fmt.Println(PI, NAME)
	fmt.Println(A, B, C)

	// %v for value of any variable of any type
	// %t for datatype of the variable
	fmt.Printf(NAME+" :%v of type %T", PI, A)

	// for more of these formatting % symbols
	// https://www.w3schools.com/go/go_formatting_verbs.php

	fmt.Printf("%b\n", C)
	fmt.Printf("%d\n", C)
	fmt.Printf("%+b\n", C)
	fmt.Printf("%o\n", C)
	fmt.Printf("%x\n", C)
	fmt.Printf("%4d\n", C)
	fmt.Printf("%05d\n", C)

}
