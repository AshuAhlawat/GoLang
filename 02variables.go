package main

import "fmt"

// var can be declared outside functions but not with WalrusOperator " := "
var e string

func main() {
	// Datatype is mentioned and that it is a Variable
	var x string = "Both var and type"
	// Datatype is inferred
	var y = 5
	// the walrus sign tells that its a variable and datatype is also inferred
	z := 3.53

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	// in Go all variables are initialized if no value is gives defaults are used
	var a string  // empty string ""
	var b int     // 0
	var c float32 // 0.0
	var d bool    // false

	c = 4.34
	e = x

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// multiple variables decleration in a single line
	y, b = b, y
	// multiple declarions must of same datatype which is given if var keyword is uses
	var f, g int = 8, 5
	h, i := "wow", true

	fmt.Println(y,b)
	fmt.Println(f,g)
	fmt.Println(h,i)

	// muliple variable declerations in a block

	var (
		j int
		k string = "amazing"
	)
	
	fmt.Println(j)
	fmt.Println(k)
}
