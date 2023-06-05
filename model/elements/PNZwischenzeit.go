package elements

import (
	"errors"
	"fmt"
	"github.com/konrad2002/dsvparser/model/types"
	"strconv"
)

type PNZwischenzeit struct {
	VeranstaltungsIdSchwimmer int
	Wettkampfnummer           int
	Wettkampfart              string
	Distanz                   int
	Zwischenzeit              types.Zeit
}

func NewPNZwischenzeit(lits []string) (PNZwischenzeit, error) {
	args := 5
	if len(lits) != args {
		return PNZwischenzeit{}, fmt.Errorf("falsche Anzahl an Argumenten für PNZWISCHENZEIT, %d statt %d", len(lits), args)
	}
	var el PNZwischenzeit
	var err1, err2, err3, err4 error
	el.VeranstaltungsIdSchwimmer, err1 = strconv.Atoi(lits[0])
	el.Wettkampfnummer, err2 = strconv.Atoi(lits[1])
	el.Wettkampfart = lits[2]
	el.Distanz, err3 = strconv.Atoi(lits[3])
	el.Zwischenzeit, err4 = types.NewZeit(lits[4])

	return el, errors.Join(err1, err2, err3, err4)
}
