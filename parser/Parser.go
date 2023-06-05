package parser

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

	//fmt.Printf("scanned: %d\n", tok)

	switch tok {
	case ABSCHNITT:
		el, err = elements.NewAbschnitt(lits)
	case AUSRICHTER:
		el, err = elements.NewAusrichter(lits)
	case AUSSCHREIBUNGIMNETZ:
		el, err = elements.NewAusschreibungImNetz(lits)
	case BANKVERBINDUNG:
		el, err = elements.NewBankverbindung(lits)
	case BESONDERES:
		el, err = elements.NewBesonderes(lits)
	case ERZEUGER:
		el, err = elements.NewErzeuger(lits)
	case FORMAT:
		el, err = elements.NewFormat(lits)
	case KAMPFGERICHT:
		el, err = elements.NewKampfgericht(lits)
	case MELDEADRESSE:
		el, err = elements.NewMeldeadresse(lits)
	case MELDEGELD:
		el, err = elements.NewMeldegeld(lits)
	case MELDESCHLUSS:
		el, err = elements.NewMeldeschluss(lits)
	case NACHWEIS:
		el, err = elements.NewNachweis(lits)
	case PFLICHTZEIT:
		el, err = elements.NewPflichtzeit(lits)
	case PNERGEBNIS:
		el, err = elements.NewPNErgebnis(lits)
	case PNREAKTION:
		el, err = elements.NewPNReaktion(lits)
	case PNZWISCHENZEIT:
		el, err = elements.NewPNZwischenzeit(lits)
	case STABLOESE, STAFFELPERSON, STZWISCHENZEIT, STERGEBNIS:
		el, err = nil, nil
	case VERANSTALTER:
		el, err = elements.NewVeranstalter(lits)
	case VERANSTALTUNG:
		el, err = elements.NewVeranstaltung(lits)
	case VERANSTALTUNGSORT:
		el, err = elements.NewVeranstaltungsort(lits)
	case VEREIN:
		el, err = elements.NewVerein(lits)
	case WERTUNG:
		el, err = elements.NewWertung(lits)
	case WETTKAMPF:
		el, err = elements.NewWettkampf(lits)
	case DATEIENDE:
		el, err = elements.NewDateiende()
	case COMMENT, EMPTY:
		el, err = nil, nil
	case EOF:
		el, err = nil, fmt.Errorf("unerwartetes dateiende")
	default:
		el = nil
		err = fmt.Errorf("unbekannter Token nach Scan: %d", tok)
	}
	return el, err
}
