package main

import (
	"fmt"
	"os"

	"github.com/krernertok/glox"
)

func runPrompt() {
	fmt.Println("got here")
}

func runFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error reading file: %q\n", path)
		fmt.Println(err)
		// EX_IOERR
		os.Exit(74)
	}
	glox.Run(file)
}

func main() {
	argLength := len(os.Args)
	if argLength > 2 {
		fmt.Println("Usage: glox [script]")
		// EX_USAGE
		os.Exit(64)
	} else if argLength == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
}
