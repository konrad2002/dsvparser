package elements

import "fmt"

type Lastschrift struct {
	Hinweis bool
}

func NewLastschrift(lits []string) (Lastschrift, error) {
	args := 1
	if len(lits) != args {
		return Lastschrift{}, fmt.Errorf("falsche Anzahl an Argumenten für LASTSCHRIFT, %d statt %d", len(lits), args)
	}
	var el Lastschrift
	el.Hinweis = lits[0] == "J"

	return el, nil
}
