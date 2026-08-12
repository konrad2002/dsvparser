package elements

import (
	"fmt"
)

type Bankverbindung struct {
	NameBank     string
	IBAN         string
	BIC          string
	Kontoinhaber string
}

func NewBankverbindung(lits []string) (Bankverbindung, error) {
	args := 3
	argsv8 := 4
	if len(lits) != args && len(lits) != argsv8 {
		return Bankverbindung{}, fmt.Errorf("falsche Anzahl an Argumenten für BANKVERBINDUNG, %d statt %d oder %d", len(lits), args, argsv8)
	}
	var el Bankverbindung
	el.NameBank = lits[0]
	el.IBAN = lits[1]
	el.BIC = lits[2]
	if len(lits) == argsv8 {
		el.Kontoinhaber = lits[3]
	}
	return el, nil
}
