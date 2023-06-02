package main

import (
	"bufio"
	"io"
)

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() {
	_ = s.r.UnreadRune()
}

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	if isSemicolon(ch) {
		s.unread()
		//return s.scanWhitespace()
	} else if isLetter(ch) {
		s.unread()
		//return s.scanIdent()
	}

	// Otherwise read the individual character.
	switch ch {
	case eof:
		return EOF, ""
	case ';':
		return SEMICOLON, string(ch)
	}

	return ILLEGAL, string(ch)
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isSemicolon(ch rune) bool {
	return ch == ';'
}

var eof = rune(0)
