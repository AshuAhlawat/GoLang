package main

import "fmt"

func main(){

	// to take n as input
	var n int
	print("Enter a number: ");fmt.Scan(&n)

	// nested for loop for the stairs pattern
	for i:=0; i<n; i++{
		for j:=0; j<i+1; j++{
			print("*")
		}
		print("\n")
	}
	
	// creating a table array
	print("\n")
	var table = [10]int{}

	// based on index changing values
	for i:=0; i<10; i++{
		table[i] = (i+1)*n
	}

	fmt.Println(table)

	// range function to print the array/list/map
	for i, el := range table{
		fmt.Println(i, el)
	}


	for i := range [10]int{}{
		print(i)
	}

}
