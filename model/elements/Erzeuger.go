package elements

import "fmt"

type Erzeuger struct {
	Software string
	Version  string
	Kontakt  string
}

func NewErzeuger(lits []string) (Erzeuger, error) {
	args := 3
	if len(lits) != args {
		return Erzeuger{}, fmt.Errorf("zu wenig Argumente für ERZEUGER, %d statt %d", len(lits), args)
	}
	var el Erzeuger
	el.Software = lits[0]
	el.Version = lits[1]
	el.Kontakt = lits[2]
	return el, nil
}
