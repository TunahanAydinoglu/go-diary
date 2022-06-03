package main

import "fmt"

func main() {
	var slc []int = []int{1, 3, 2, 3, 3, 2, 6, 7, 4, 5, 6}

	fmt.Println("input", slc)

	distinctResponse := distinct(slc)
	fmt.Println("output", distinctResponse)
}

func distinct(slc []int) []int {
	copy := make([]int, 0)
	keys := make(map[int]int)
	for i, v := range slc {
		if _, ok := keys[v]; !ok {
			keys[v] = i
			copy = append(copy, v)
		}
	}

	return copy
}
