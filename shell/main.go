package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ") // Prompt
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		// Execute
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil
	}

	args := strings.Split(input, " ")
	if args[0] == "exit" {
		os.Exit(0)
	}

	// This translates to: exec.Command("cmd", "/C", "user_command", "arg1", "arg2", ...)
	cmd := exec.Command("cmd", append([]string{"/C"}, args...)...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
