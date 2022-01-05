package main

import "fmt"

func main() {

	array := [...]int{12, 24, 34}
	slice := array[:]
	fmt.Println(slice)

	slice[0] = 10
	fmt.Println(slice)
	//ref type
	fmt.Println(array)

	var langs = []string{"go", "c#", "java"}
	langs = append(langs, "javascript")
	fmt.Println(langs)

	langs = append(langs[1:len(langs)]) //remove first item
	fmt.Println(langs)

	langs = append(langs[:len(langs)-1]) //remove last item
	fmt.Println(langs)

	numbers := make([]int, 5, 5)

	for i := 0; i < len(numbers); i++ {
		numbers[i] = 3 + i*7
	}

	fmt.Println(numbers)

}
