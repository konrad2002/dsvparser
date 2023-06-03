package elements

import "fmt"

type AusschreibungImNetz struct {
	Internetadresse string
}

func NewAusschreibungImNetz(lits []string) (AusschreibungImNetz, error) {
	args := 1
	if len(lits) != args {
		return AusschreibungImNetz{}, fmt.Errorf("ungütlige Attribute für AUSSCHREIBUNGIMNETZ, %d statt %d", len(lits), args)
	}
	return AusschreibungImNetz{Internetadresse: lits[0]}, nil
}
