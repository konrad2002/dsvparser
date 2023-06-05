package elements

import (
	"errors"
	"fmt"
	"strconv"
)

type Verein struct {
	Vereinsbezeichnung   string
	Vereinskennzahl      int
	Landesschwimmverband int
	FinaNationenkuerzel  string
}

func NewVerein(lits []string) (Verein, error) {
	args := 4
	if len(lits) != args {
		return Verein{}, fmt.Errorf("falsche Anzahl an Argumenten für VEREIN, %d statt %d", len(lits), args)
	}
	var el Verein
	var err1, err2 error
	el.Vereinsbezeichnung = lits[0]
	el.Vereinskennzahl, err1 = strconv.Atoi(lits[1])
	el.Landesschwimmverband, err2 = strconv.Atoi(lits[2])
	el.FinaNationenkuerzel = lits[3]
	return el, errors.Join(err1, err2)
}
