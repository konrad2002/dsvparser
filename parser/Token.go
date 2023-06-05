package parser

type Token int

const (
	ILLEGAL Token = iota
	EOF
	COMMENT
	EMPTY

	ABSCHNITT
	AUSRICHTER
	AUSSCHREIBUNGIMNETZ
	BANKVERBINDUNG
	BESONDERES
	DATEIENDE
	ERZEUGER
	FORMAT // D,E
	KAMPFGERICHT
	MELDEADRESSE
	MELDEGELD
	MELDESCHLUSS
	NACHWEIS
	PFLICHTZEIT
	PNERGEBNIS
	PNREAKTION
	PNZWISCHENZEIT
	STABLOESE
	STAFFELPERSON
	STERGEBNIS
	STZWISCHENZEIT
	VERANSTALTER
	VERANSTALTUNG
	VERANSTALTUNGSORT
	VEREIN
	WERTUNG
	WETTKAMPF
)

func NewToken(value string) (tok Token) {
	switch value {
	case "FORMAT":
		tok = FORMAT
	case "ERZEUGER":
		tok = ERZEUGER
	case "VERANSTALTUNG":
		tok = VERANSTALTUNG
	case "VERANSTALTUNGSORT":
		tok = VERANSTALTUNGSORT
	case "AUSSCHREIBUNGIMNETZ":
		tok = AUSSCHREIBUNGIMNETZ
	case "VERANSTALTER":
		tok = VERANSTALTER
	case "AUSRICHTER":
		tok = AUSRICHTER
	case "MELDEADRESSE":
		tok = MELDEADRESSE
	case "MELDESCHLUSS":
		tok = MELDESCHLUSS
	case "BANKVERBINDUNG":
		tok = BANKVERBINDUNG
	case "BESONDERES":
		tok = BESONDERES
	case "NACHWEIS":
		tok = NACHWEIS
	case "ABSCHNITT":
		tok = ABSCHNITT
	case "WETTKAMPF":
		tok = WETTKAMPF
	case "WERTUNG":
		tok = WERTUNG
	case "PFLICHTZEIT":
		tok = PFLICHTZEIT
	case "MELDEGELD":
		tok = MELDEGELD
	case "DATEIENDE":
		tok = DATEIENDE
	case "KAMPFGERICHT":
		tok = KAMPFGERICHT
	case "VEREIN":
		tok = VEREIN
	case "PNERGEBNIS":
		tok = PNERGEBNIS
	case "PNZWISCHENZEIT":
		tok = PNZWISCHENZEIT
	case "PNREAKTION":
		tok = PNREAKTION
	case "STERGEBNIS":
		tok = STERGEBNIS
	case "STAFFELPERSON":
		tok = STAFFELPERSON
	case "STZWISCHENZEIT":
		tok = STZWISCHENZEIT
	case "STABLOESE":
		tok = STABLOESE
	default:
		tok = ILLEGAL
	}
	return tok
}