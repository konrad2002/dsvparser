package elements

import "fmt"

type Veranstaltung struct {
	Veranstaltungsbezeichnung string
	Veranstaltungsort         string
	Bahnlaenge                string
	Zeitmessung               string
}

func NewVeranstaltung(lits []string) (Veranstaltung, error) {
	args := 4
	if len(lits) != args {
		return Veranstaltung{}, fmt.Errorf("zu wenig Argumente für VERANSTALTUNG, %d statt %d", len(lits), args)
	}
	var el Veranstaltung
	el.Veranstaltungsbezeichnung = lits[0]
	el.Veranstaltungsort = lits[1]
	el.Bahnlaenge = lits[2]
	el.Zeitmessung = lits[3]
	return el, nil
}
