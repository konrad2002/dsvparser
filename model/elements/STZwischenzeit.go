package elements

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/konrad2002/dsvparser/model/types"
)

type STZwischenzeit struct {
	VeranstaltungsIdStaffel     int
	Wettkampfnummer             int
	Wettkampfart                string
	StartnummerSchwimmerStaffel int
	Distanz                     int
	Zwischenzeit                types.Zeit
}

func NewSTZwischenzeit(lits []string) (STZwischenzeit, error) {
	args := 6
	if len(lits) != args {
		return STZwischenzeit{}, fmt.Errorf("falsche Anzahl an Argumenten für STZWISCHENZEIT, %d statt %d", len(lits), args)
	}
	var el STZwischenzeit
	var err1, err2, err3, err4, err5 error
	el.VeranstaltungsIdStaffel, err1 = strconv.Atoi(lits[0])
	el.Wettkampfnummer, err2 = strconv.Atoi(lits[1])
	el.Wettkampfart = lits[2]
	el.StartnummerSchwimmerStaffel, err3 = strconv.Atoi(lits[3])
	el.Distanz, err4 = strconv.Atoi(lits[4])
	el.Zwischenzeit, err5 = types.NewZeit(lits[5])

	return el, errors.Join(err1, err2, err3, err4, err5)
}
