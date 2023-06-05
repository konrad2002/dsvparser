package elements

import "fmt"

type Besonderes struct {
	Anmerkungen string
}

func NewBesonderes(lits []string) (Besonderes, error) {
	args := 1
	if len(lits) != args {
		return Besonderes{}, fmt.Errorf("falsche Anzahl an Argumenten für BESONDERES, %d statt %d", len(lits), args)
	}
	var el Besonderes
	el.Anmerkungen = lits[0]
	return el, nil
}
