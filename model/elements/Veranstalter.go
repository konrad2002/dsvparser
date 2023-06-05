package elements

import "fmt"

type Veranstalter struct {
	Name string
}

func NewVeranstalter(lits []string) (Veranstalter, error) {
	args := 1
	if len(lits) != args {
		return Veranstalter{}, fmt.Errorf("falsche Anzahl an Argumenten für VERANSTALTER, %d statt %d", len(lits), args)
	}
	var el Veranstalter
	el.Name = lits[0]
	return el, nil
}
