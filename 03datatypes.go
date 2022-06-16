package main

import "fmt"

func main() {
	// equal to your default system achitecture(32, 64) bits
	var a int = -45
	// can only be positive
	// from 0 to int upper limit - int lower limit
	var b uint = 32
	// from -128 to 127
	var c int8 = 5
	// from about -32K to 32K
	var d int16 = 1234
	// from about -2Billion to +2Billion
	var e int32 = 4121212
	// from -9.2233e+18 to +92233e+18
	var f int64 = 72391231923988

	fmt.Println(a, b, c, d, e, f)

	
	// from -3.4e+38 to 3.4e+38
	var g float32 = 5.2e+10
	// from -1.7e+308 to +1.7e+308
	var h float64 = 2.412e+28

	fmt.Println(g, h)

}
