package main

import (
	"bufio"
	"bytes"
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

// Scan returns the next token and all following literal values.
func (s *Scanner) Scan() (tok Token, lit []string) {

	lc := ' '
	ch := ' '
	var buf bytes.Buffer

	// read token
	for {
		if ch = s.read(); ch == eof {
			return EOF, nil
		}
		if ch == ':' {
			break
		}
		if ch == '*' && lc == '(' {
			for {
				if ch = s.read(); ch == eof || ch == '\n' {
					return COMMENT, nil
				}
			}
		}

		buf.WriteRune(ch)
		lc = ch
	}

	tok = NewToken(buf.String())

	buf.Reset()

	// read attributes
	for {
		if ch = s.read(); ch == eof {
			break
		}
		if ch == '\n' {
			break
		}

		if isSemicolon(ch) {
			lit = append(lit, buf.String())
			buf.Reset()
			continue
		}

		buf.WriteRune(ch)

	}

	return tok, lit
}

func isSemicolon(ch rune) bool {
	return ch == ';'
}

var eof = rune(0)
