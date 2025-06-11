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
	if arg[0] == arg[1] {
		fmt.Println("file can not be same")
		return
	}
	if !strings.HasSuffix(arg[0], ".txt") {
		fmt.Println("input file is not a .txt file")
		return
	}

	if !strings.HasSuffix(arg[1], ".txt") {
		fmt.Println("output file is not a .txt file")
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
		fmt.Println("error writing file")
		return
	}
}
