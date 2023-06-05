package elements

import (
	"errors"
	"fmt"
	"github.com/konrad2002/dsvparser/model/types"
	"strconv"
)

type Wertung struct {
	Wettkampfnummer int
	Wettkampfart    rune
	WertungsID      int
	Wertungsklasse  types.Wertungsklasse
	MindestJahrgang string
	MaximalJahrgang string
	Geschlecht      types.Geschlecht
	Wertungsname    string
}

func NewWertung(lits []string) (Wertung, error) {
	args := 8
	if len(lits) != args {
		return Wertung{}, fmt.Errorf("falsche Anzahl an Argumenten für WERTUNG, %d statt %d", len(lits), args)
	}
	var el Wertung
	var err1, err2, err3, err4 error
	el.Wettkampfnummer, err1 = strconv.Atoi(lits[0])
	if len(lits[1]) == 1 {
		el.Wettkampfart = []rune(lits[1])[0]
	} else {
		err2 = fmt.Errorf("wettkampfart ist kein Zeichen")
	}
	el.WertungsID, err3 = strconv.Atoi(lits[2])
	el.Wertungsklasse = types.NewWertungsklasse(lits[3])
	el.MindestJahrgang = lits[4]
	el.MaximalJahrgang = lits[5]
	if len(lits[6]) > 0 {
		el.Geschlecht = types.NewGeschlecht([]rune(lits[6])[0])
	}
	el.Wertungsname = lits[7]
	return el, errors.Join(err1, err2, err3, err4)
}
