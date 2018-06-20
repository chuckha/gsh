package gsh

import (
	"bufio"
	"io"
)

type Scanner struct {
	scanner *bufio.Scanner
}

func NewScanner(input io.Reader) *Scanner {
	scanner := bufio.NewScanner(input)
	return &Scanner{
		scanner: scanner,
	}
}

func (s *Scanner) Input() (string, error) {
	for s.scanner.Scan() {
		return s.scanner.Text(), s.scanner.Err()
	}
	return "", nil
}
