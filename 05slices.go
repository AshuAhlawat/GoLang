package main

import "fmt"

// The alternate of python lists, slices are just dynamic lists
func main(){
	// simple interpretted length array
	var arr = [...]int{1,2,3,4,5}

	// taking out a slice of the arr
	var test = arr[2:4]

	fmt.Printf("%v is a Go slice of length %v from the array %v",test,len(test),arr)

	test[0] = 5
	// first param is the slice and it returns a slice with other params appended to its end
	test = append(test, arr[0], arr[1])

	fmt.Printf("\n%v is our slice\n",test)

	// extending slices from slices / only one at a time 
	test = append(test, test[0:2]...)

	fmt.Printf("Our final slice %v of length %v\n\n",test, len(test))

	// as go loads all the items to slice into memory so if the slice is large thats too much data
	// instead we can copy the items we need into a new array

	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	fmt.Printf("Our new slice\n%v\n", numbers)

	// creating slice of the needed data
	needed_nums := numbers[0:len(numbers)/2]
	fmt.Println(needed_nums)

	// new empty array of length equal to our slice
	new_numbers := make([]int, len(needed_nums))
	fmt.Println(new_numbers)

	// copying slice data into empty array
	copy(new_numbers, needed_nums)
	fmt.Println(new_numbers)

}