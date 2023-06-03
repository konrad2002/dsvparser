package main

import (
	"fmt"
	"github.com/konrad2002/dsvparser/model/elements"
	"io"
)

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

// Parse receives token and literal from scanner and wraps it inside element model
func (p *Parser) Parse() (el interface{}, err error) {
	tok, lits := p.s.Scan()
	if tok == ILLEGAL {
		return nil, fmt.Errorf("scanned illegal token")
	}

	switch tok {
	case ABSCHNITT:
		el, err = elements.NewAbschnitt(lits)
	case AUSRICHTER:
		el, err = elements.NewAusrichter(lits)
		// TODO fehlende Elemente
	default:
		el = nil
		err = fmt.Errorf("unbekannter Token nach Scan")
	}
	return el, err
}
