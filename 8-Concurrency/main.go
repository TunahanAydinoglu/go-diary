package main

import (
	"fmt"
	"time"
)

func main() {
	callGoRoutineFunc()
	fmt.Println("---------------------------------------")
	callParallelRoutineFunc()
}

func goroutineExFunc() {
	for index := byte('a'); index <= byte('f'); index++ {
		fmt.Println(string(index))
	}
}

func callGoRoutineFunc() {
	go goroutineExFunc()
	fmt.Println("Hi") //is printed first
	go fmt.Println("Say hello")
	time.Sleep(100 * time.Millisecond)
}

func parallelExFunc() {
	for index := byte('a'); index <= byte('f'); index++ {
		go fmt.Println(string(index))
	}
}
func callParallelRoutineFunc() {
	go parallelExFunc()
	fmt.Println("Hi")
	go fmt.Println("Say hello")
	time.Sleep(100 * time.Millisecond)
}
