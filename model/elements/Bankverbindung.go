package elements

import (
	"fmt"
)

type Bankverbindung struct {
	NameBank string
	IBAN     string
	BIC      string
}

func NewBankverbindung(lits []string) (Bankverbindung, error) {
	args := 3
	if len(lits) != args {
		return Bankverbindung{}, fmt.Errorf("zu wenig Argumente für BANKVERBINDUNG, %d statt %d", len(lits), args)
	}
	var el Bankverbindung
	el.NameBank = lits[0]
	el.IBAN = lits[1]
	el.BIC = lits[2]
	return el, nil
}
