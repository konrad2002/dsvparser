package elements

import "fmt"

type Meldeadresse struct {
	Name    string
	Strasse string
	PLZ     string
	Ort     string
	Land    string
	Telefon string
	Fax     string
	Email   string
}

func NewMeldeadresse(lits []string) (Meldeadresse, error) {
	args := 8
	if len(lits) != args {
		return Meldeadresse{}, fmt.Errorf("zu wenig Argumente für MELDEADRESSE, %d statt %d", len(lits), args)
	}
	var el Meldeadresse
	el.Name = lits[0]
	el.Strasse = lits[1]
	el.PLZ = lits[2]
	el.Ort = lits[3]
	el.Land = lits[4]
	el.Telefon = lits[5]
	el.Fax = lits[6]
	el.Email = lits[7]
	return el, nil
}
