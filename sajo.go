package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/anshulkanwar/sajo/internal"
)

var hadError = false

func runFile(path string) error {
	bytes, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	run(string(bytes))
	if hadError {
		os.Exit(65)
	}
	return nil
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		run(line)
		hadError = false
	}
}

func run(source string) {
	scanner := internal.NewScanner(source)
	tokens := scanner.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Usage: sajo [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}
