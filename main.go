package main

import (
	"fmt"
	"os"
	cleaner "reload/cleaner"
	command "reload/commands"
	"strings"
)

func procedures(s string) string {
	res := cleaner.Clean_text(s)
	res = command.Proceed_commands(res)
	res = cleaner.Clean_text(res)
	res = strings.TrimSpace(res)
	return res
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		fmt.Println("there must be exactly 2 inputs (Usage: go run . sample.txt output.txt)")
		return
	}
	bytes, err := os.ReadFile(arg[0])
	if err != nil {
		fmt.Println("error reading input file")
		return
	}
	fileText := string(bytes[:])
	res := procedures(fileText)
	err = os.WriteFile(arg[1], []byte(res), 0o666)
	if err != nil {
		return
	}
}
