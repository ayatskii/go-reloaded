package main

import (
	"os"
	cleaner "reload/cleaner"
	command "reload/commands"
	"strings"
)

func main() {
	arg := os.Args[1:]
	bytes, err := os.ReadFile(arg[0])
	if err != nil {
		return
	}
	fileText := string(bytes[:])
	res := cleaner.Clean_text(fileText)
	res = command.Proceed_commands(res)
	res = cleaner.Clean_text(res)
	res = strings.TrimSpace(res)
	err = os.WriteFile(arg[1], []byte(res), 0666)
	if err != nil {
		return
	}
}
