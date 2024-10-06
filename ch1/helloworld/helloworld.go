package main

import (
	"fmt"
	"os"
)

func main() {
	const greeting = "Hello"
	fmt.Printf("%s, World1!\n", greeting)

	fmt.Println(getArgs())
}

func getArgs() []string {
	return os.Args // [./helloworld argument1]
	// return os.Args[1:2] // [argument1]
}
