package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/chuckha/gsh"
)

// Stdin and stdout are all sorts of fucked.

type Prompter interface {
	Prompt()
}

type InputGetter interface {
	Input() (string, error)
}

type Parser interface {
	Parse(string) ([]string, error)
}

type emptyerror interface {
	IsEmpty() bool
}

func main() {
	// Setting up dependencies
	var prompter Prompter
	prompter = &gsh.Prompter{
		String: "$ ",
	}
	var inputGetter InputGetter
	inputGetter = gsh.NewScanner(os.Stdin)
	var parser Parser
	parser = &gsh.Parser{}

	fmt.Println("Welcome to gsh! It's bash wrapped up in go for whatever reason")
	for {
		prompter.Prompt()
		input, err := inputGetter.Input()
		if err != nil {
			fmt.Printf("error scanning text: %v", err)
			continue
		}
		args, err := parser.Parse(input)
		if err != nil {
			switch err.(type) {
			case emptyerror:
			default:
				fmt.Printf("error parsing args: %v\n", err)
			}
			continue
		}
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}
}

// something prints the prompt
// something reads the user input
// something parses the user input
// something execs the parsed input
// something handles the file descriptors
