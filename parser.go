package main

import "io"

type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// should be called for a single line inside a dsv file
func (p *Parser) Parse() (interface{}, error) {
	return nil, nil
}
