package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	c := make(chan int)
	go sumFunc(numbers[:len(numbers)], c)
	go sumFunc(numbers[:len(numbers)/2], c)
	x, y := <-c, <-c

	fmt.Println(x, y)

}

func sumFunc(numbers []int, c chan int) {
	sum := 0

	for _, value := range numbers {
		sum += value
	}
	c <- sum
}
