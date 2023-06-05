package elements

import (
	"fmt"
	"strconv"
)

type Meldegeld struct {
	MeldegeldTyp    string
	Betrag          string
	Wettkampfnummer int
}

func NewMeldegeld(lits []string) (Meldegeld, error) {
	args := 3
	if len(lits) != args {
		return Meldegeld{}, fmt.Errorf("zu wenig Argumente für MELDEGELD, %d statt %d", len(lits), args)
	}
	var el Meldegeld
	el.MeldegeldTyp = lits[0]
	el.Betrag = lits[1]
	el.Wettkampfnummer, _ = strconv.Atoi(lits[2])
	return el, nil
}
