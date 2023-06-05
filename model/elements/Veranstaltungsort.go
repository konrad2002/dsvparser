package elements

import "fmt"

type Veranstaltungsort struct {
	NameSchwimmhalle string
	Strasse          string
	PLZ              string
	Ort              string
	Land             string
	Telefon          string
	Fax              string
	Email            string
}

func NewVeranstaltungsort(lits []string) (Veranstaltungsort, error) {
	args := 8
	if len(lits) != args {
		return Veranstaltungsort{}, fmt.Errorf("zu wenig Argumente für VERANSTALTUNGSORT, %d statt %d", len(lits), args)
	}
	var el Veranstaltungsort
	el.NameSchwimmhalle = lits[0]
	el.Strasse = lits[1]
	el.PLZ = lits[2]
	el.Ort = lits[3]
	el.Land = lits[4]
	el.Telefon = lits[5]
	el.Fax = lits[6]
	el.Email = lits[7]
	return el, nil
}
