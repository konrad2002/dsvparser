package elements

import "fmt"

type Ausrichter struct {
	NameDesAusrichters string
	Name               string
	Strasse            string
	PLZ                string
	Ort                string
	Land               string
	Telefon            string
	Fax                string
	Email              string
}

func NewAusrichter(lits []string) (Ausrichter, error) {
	args := 9
	if len(lits) != args {
		return Ausrichter{}, fmt.Errorf("ungütlige Attribute für AUSRICHTER, %d statt %d", len(lits), args)
	}
	var el Ausrichter
	el.NameDesAusrichters = lits[0]
	el.Name = lits[1]
	el.Strasse = lits[2]
	el.PLZ = lits[3]
	el.Ort = lits[4]
	el.Land = lits[5]
	el.Telefon = lits[6]
	el.Fax = lits[7]
	el.Email = lits[8]
	return el, nil
}
