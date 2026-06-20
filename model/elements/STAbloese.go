package elements

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/konrad2002/dsvparser/model/types"
)

type STAbloese struct {
	VeranstaltungsIdStaffel     int
	Wettkampfnummer             int
	Wettkampfart                string
	StartnummerSchwimmerStaffel int
	Art                         rune
	Reaktionszeit               types.Zeit
}

func NewSTAbloese(lits []string) (STAbloese, error) {
	args := 6
	if len(lits) != args {
		return STAbloese{}, fmt.Errorf("falsche Anzahl an Argumenten für STABLOESE, %d statt %d", len(lits), args)
	}
	var el STAbloese
	var err1, err2, err3, err4 error
	el.VeranstaltungsIdStaffel, err1 = strconv.Atoi(lits[0])
	el.Wettkampfnummer, err2 = strconv.Atoi(lits[1])
	el.Wettkampfart = lits[2]
	el.StartnummerSchwimmerStaffel, err3 = strconv.Atoi(lits[3])
	if len(lits[4]) > 0 {
		el.Art = []rune(lits[4])[0]
	}
	el.Reaktionszeit, err4 = types.NewZeit(lits[5])

	return el, errors.Join(err1, err2, err3, err4)
}
