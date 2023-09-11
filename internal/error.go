package internal

import "fmt"

func error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Printf("[line %d] Error %v: %v\n", line, where, message)
}
