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
			if err.Error() == "EOF" {
				fmt.Printf("\nQuitting...\n")
				os.Exit(0)
			}
			fmt.Print("Couldn't process input. Try again!\nError:\n%s\n", err)
			continue
		}

		err = glox.Run(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
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

	err = glox.Run(string(file))
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
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
