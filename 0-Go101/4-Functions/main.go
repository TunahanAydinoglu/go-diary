package main

import (
	"fmt"
)

func main() {

	name := "CLown"
	greetWithNameVoidFunc(name)
	fmt.Println(name)

	sayNameWithPointer(&name)
	fmt.Println(name)

	sum := sumMethod(13, 17)
	print(sum)
}

func greetWithNameVoidFunc(name string) {
	fmt.Println("Hi", name)
	name = "Sampyy"
}

func sayNameWithPointer(name *string) {

	fmt.Println(*name)
	*name = "Sampyy"
}

func sumMethod(x, y int) int {
	return x + y
}
