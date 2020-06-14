package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Printf("All arguments: %v\n", args)

	argsWithoutProgram := os.Args[1:]
	fmt.Printf("Arguments without program name: %v\n", argsWithoutProgram)

	for i, arg := range argsWithoutProgram {
		fmt.Printf("[%d]: %s\n", i, arg)
	}
}
