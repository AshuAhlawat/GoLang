package main

import "fmt"

// main.Person with 4 attributes
type Person struct {
  name string
  age int
  job string
  salary int
}

// other structs using structs 
type Pilot struct {
	p Person
	plane string
}

func main() {
  // declaring 2 persons and 1 pilot
  var p1 Person
  var p2 Person
  var p3 Pilot

  // p1 specification
  p1.name = "Hege"
  p1.age = 45
  p1.job = "Teacher"
  p1.salary = 6000

  // p2 specification
  p2.name = "Cecilie"
  p2.age = 24
  p2.job = "Marketing"
  p2.salary = 4500

  // Print Pers1 info by calling a function
  printPerson(p1)
  // Print Pers2 info by fmt print
  fmt.Println(p2)

  
  
  p3.p = p1
  p3.plane = "Boeingg-737"

  fmt.Printf("\n%T : %v \nFlying : %v",p3.p, p3.p, p3.plane)
}

func printPerson(pers Person) {
  fmt.Println("Name: ", pers.name)
  fmt.Println("Age : ", pers.age)
  fmt.Println("Job : ", pers.job)
  fmt.Println("Sal : ", pers.salary)
  print("\n")
}