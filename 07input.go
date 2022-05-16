package main

import "fmt"
import "math"

func main(){
	println("Calculator")

	var x,y float32

	// taking inputs of whole line stops at \n
	print("x = ")
	fmt.Scanln(&x)
	print("y = ")
	fmt.Scanln(&y)
	// same as
	// fmt.Scanf("x = %v\ny = %v",&x, &y)
	
	// printing out the operations
	fmt.Printf("\nx + y = %v", x+y)
	fmt.Printf("\nx - y = %v", x-y)
	fmt.Printf("\nx * y = %v", x*y)
	fmt.Printf("\nx / y = %v", x/y)
	fmt.Printf("\nx ^ y = %v", math.Pow(float64(x),float64(y)))

	var a,b float32

	print("\n\nEnter a , b seperated by a space : ")
	fmt.Scan(&a, &b)
	// same as
	// fmt.Scanf("%v %v",&a, &b)
	
	print("\n a > b is ",a>b)
	print("\n a == b is ",a==b)
	print("\n a < b is ",a<b)



}