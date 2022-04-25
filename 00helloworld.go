// every .go file needs to belong to a package
package main

// default module with basic functions like print
import "fmt"

// main function which gets executed first
func main() {
	// Automatic ; whenever a \n is encounted
	fmt.Print("Hello World!")
	fmt.Println("Hello World!")
	fmt.Print("Hello World!")
	fmt.Printf("Hello %s!", "World")
}
