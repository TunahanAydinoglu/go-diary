package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numberString string = "5"
	var number int = 15
	var floatNumber float32 = 3.0

	fmt.Println(numberString, number, floatNumber)

	// convert string to int
	stringToInt, _ := strconv.Atoi(numberString)
	fmt.Println(stringToInt + 2)

	//convert int to string
	intToString := strconv.Itoa(number)
	fmt.Println("Number : " + intToString)

	//Casting
	var intToFloat float64 = float64(number)
	fmt.Println(intToFloat)
}
