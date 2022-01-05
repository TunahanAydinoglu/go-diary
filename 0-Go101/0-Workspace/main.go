package main

import "fmt"

func main() {

	// pointerExp()

	iotaTryFunc()
}

func pointerExp() {

	name := "Tuna"

	var pstring *string = new(string)

	*pstring = name

	var newPoint *string = pstring

	fmt.Println(name, pstring, *pstring, &name, newPoint)
}

func iotaTryFunc() {

	const (
		first  = iota
		second = 2 << iota
		third
	)

	fmt.Println(first, second, third)
}
