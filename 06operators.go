package main

import "fmt"


func main(){
	// declaring some variables together
	var (
		a = 3
		b = 10
		c = 2.34
		d = 12.41
	)
	fmt.Printf("a = %v (int)\nb = %v\nc = %v (float64)\nd = %v\n\n",a,b,c,d)


	// Arithematic
	fmt.Println("Arithematic Operations")
	// can only operate on similar datatypes
	x := a+b
	y := d-c
	fmt.Printf("a+b = %v (%T) \nd-c = %v (%T)\n\n", x,x, y,y)


	// Assignment
	fmt.Println("Assignment Operations")
	// float only divided by float and gives exact division
	z := d/c
	print("d/c + 3 - 10 = ",z)
	z++;z++;z++
	print(" + 3 - 10 = ",z," - 10")
	z -= 10
	print(" = ",z," (float64)")
	
	// int / int is always the quoutient of them
	e := b/a
	fmt.Printf("\nb/a = %v (%T)", e,e)
	e = b%a
	fmt.Printf("\nb%%a = %v (%T)", e,e)


	// Comparison
	fmt.Println("\n\nComparison Operations")
	// different data types cant even be compared
	print("e>a = ",e>a, " (boolean)")
	print("\nc!=z = ",c!=z)


	// Logical
	fmt.Println("\n\nLogical Operations")
	print("f = a>b and c<d = ", a<b, " and ",c>d, " = ", a<b && c>d)
	f := false
	print("\n!f = !",f ," = ", !f)

}