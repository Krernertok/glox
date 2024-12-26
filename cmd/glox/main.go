package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/krernertok/glox"
)

func runPrompt() {
	input := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("> ")

		input, err := input.ReadString('\n')
		if err != nil {
			fmt.Println("Couldn't process input. Try again!")
			continue
		}

		glox.Run(input)
	}
}

func runFile(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file: %q\n", path)
		fmt.Println(err)
		// EX_IOERR
		os.Exit(74)
	}

	glox.Run(string(file))
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
