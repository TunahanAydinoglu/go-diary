package main

import (
	"fmt"
	"os"
)

func main() {
	// getEnvironment()

	uName := os.Getenv("USERNAME")
	uDomain := os.Getenv("USERDOMAIN")
	processorArchitecture := os.Getenv("PROCESSOR_ARHITECTURE")
	processorIdentifier := os.Getenv("PROCESSOR_IDENTIFIER")
	processorLevel := os.Getenv("PROCESSOR_LEVEL")
	goRoot := os.Getenv("GOROOT")
	goPath := os.Getenv("GOPATH")
	homePath := os.Getenv("HOMEPATH")

	fmt.Println("Username : " + uName)
	fmt.Println("Domain : " + uDomain)
	fmt.Println("Processor Architecture : " + processorArchitecture)
	fmt.Println("Processor Identifier : " + processorIdentifier)
	fmt.Println("Processor Level : " + processorLevel)
	fmt.Println("Go Root : " + goRoot)
	fmt.Println("Go Path : " + goPath)
	fmt.Println("Home Path : " + homePath)
}

func getEnvironment() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}
