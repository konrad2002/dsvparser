package elements

import (
	"errors"
	"fmt"
	"github.com/konrad2002/dsvparser/model/types"
	"strconv"
)

type PNReaktion struct {
	VeranstaltungsIdSchwimmer int
	Wettkampfnummer           int
	Wettkampfart              string
	Art                       rune
	Reaktionszeit             types.Zeit
}

func NewPNReaktion(lits []string) (PNReaktion, error) {
	args := 5
	if len(lits) != args {
		return PNReaktion{}, fmt.Errorf("falsche Anzahl an Argumenten für PNREAKTION, %d statt %d", len(lits), args)
	}
	var el PNReaktion
	var err1, err2, err3 error
	el.VeranstaltungsIdSchwimmer, err1 = strconv.Atoi(lits[0])
	el.Wettkampfnummer, err2 = strconv.Atoi(lits[1])
	el.Wettkampfart = lits[2]
	if len(lits[3]) > 0 {
		el.Art = []rune(lits[3])[0]
	}
	el.Reaktionszeit, err3 = types.NewZeit(lits[4])

	return el, errors.Join(err1, err2, err3)
}
