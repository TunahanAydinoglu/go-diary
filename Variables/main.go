package main

import "fmt"

// number:=5 //not working in scope
var number = 5

var (
	value1 = "value1"
	value2 = "value2"
)

func main() {
	var message string
	message = "Hello Clown!"
	fmt.Println(message)

	var name, surname string = "Tuna", "Aydinoglu"
	fmt.Println(name, surname)

	var value = "Tuna"
	// value = 5 // auto type casting not working
	fmt.Println(value)

	age := 24
	fmt.Println(age)

	b := "string"
	b = `deneme`
	fmt.Println(b)

	fmt.Println(number)

	const description = "its a constant"
	// description = "its not changes"
	fmt.Println(description)

	fmt.Println(value1)
}
