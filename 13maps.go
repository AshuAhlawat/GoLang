package main

import "fmt"

func main() {
	// declaring initial pairs
    b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}
	
	// to delete key value pairs
	delete(b, "Trondheim")
    fmt.Printf("b  %v\n\n", b)
	
	// to add pairs 1 by 1
    var a = make(map[string]string) // The map is empty now
    a["brand"] = "Ford"
    a["model"] = "Mazda"
    a["year"] = "1964"

	a["model"] = "Mustang"

	fmt.Printf("Car : %v %v\nYear : %v",a["brand"], a["model"], a["year"])
  	
}


// keys allowed

// Booleans
// Numbers
// Strings
// Arrays
// Pointers
// Structs