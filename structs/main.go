package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var alex person // defaults to "" for each string property

	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
