package main

import "fmt"

func main() {
	// size is declared and extra spaces are filled with default values for that datatype
	var arr1 = [4]int{2, 3}
	arr1[3] = 1

	// without var keyword and length is inferred
	arr2 := [...]string{"hello", "dear"}
	
	// only initializing certain elements
	var arr3 = [...]float32{2:3.212, 6:5.8704}

	fmt.Println(arr1, len(arr1), arr1[1])
	fmt.Println(arr2, len(arr2), arr2[1])
	fmt.Println(arr3, len(arr3), arr3[1])
}
