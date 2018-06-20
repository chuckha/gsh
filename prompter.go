package gsh

import "fmt"

type Prompter struct {
	String string
}

func (p *Prompter) Prompt() {
	fmt.Print(p.String)
}
