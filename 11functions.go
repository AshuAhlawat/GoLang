package main

import "fmt"

// simple function that returns nothing
func introduction(name string, age int, hobbies []string){
	fmt.Printf("Myself %v, I am %v years old, and i like ", name, age)
	for i := range hobbies{
		print(hobbies[i],", ")
	}
	print("etc.")
}

// function returning int
func add(x int, y int) int{
	return x+y
}

// function returning 2 values
func muldiv(x int, y int) (int,float32) {
	
	return x*y, float32(x)/float32(y)
}

func main(){
	likes := []string{"watching anime", "gaming"}
	name := "Ashu"
	age := 20

	introduction(name, age, likes)

	print("\n",add(7,add(3,9)),"\n")

	mul, div := muldiv(5,4)

	print("x*y = ", mul, "\n")
	print("x/y = ", div, "\n")

}