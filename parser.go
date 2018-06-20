package gsh

import "strings"

type Parser struct{}

func (p *Parser) Parse(input string) ([]string, error) {
	if input == "" {
		return nil, NewEmptyInputError()
	}
	return strings.Split(input, " "), nil
}

type EmptyInput struct {
	Err error
}

func NewEmptyInputError() error {
	return &EmptyInput{}
}
func (e *EmptyInput) Error() string {
	return "Input was empty"
}
func (e *EmptyInput) IsEmpty() bool {
	return true
}
