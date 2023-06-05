package main

import (
	"fmt"
	"github.com/konrad2002/dsvparser/model/elements"
	"io"
)

type Parser struct {
	s *Scanner
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
	case AUSSCHREIBUNGIMNETZ:
		el, err = elements.NewAusschreibungImNetz(lits)
	case FORMAT:
		el, err = elements.NewFormat(lits)
	// TODO: fehlende Elemente
	case DATEIENDE:
		el, err = elements.NewDateiende()
	case COMMENT:
		el, err = nil, nil
	case EOF:
		el, err = nil, fmt.Errorf("unerwartetes dateiende")
	default:
		el = nil
		err = fmt.Errorf("unbekannter Token nach Scan")
	}
	return el, err
}
