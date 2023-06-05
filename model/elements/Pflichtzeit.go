package elements

import (
	"errors"
	"fmt"
	"github.com/konrad2002/dsvparser/model/types"
	"strconv"
)

type Pflichtzeit struct {
	Wettkampfnummer int
	Wettkampfart    rune
	Wertungsklasse  types.Wertungsklasse
	MindestJahrgang string
	MaximalJahrgang string
	Pflichtzeit     types.Zeit
	Geschlecht      types.Geschlecht
}

func NewPflichtzeit(lits []string) (Pflichtzeit, error) {
	args := 7
	if len(lits) != args {
		return Pflichtzeit{}, fmt.Errorf("falsche Anzahl an Argumenten für PFLICHTZEIT, %d statt %d", len(lits), args)
	}
	var el Pflichtzeit
	var err1, err2, err3 error
	el.Wettkampfnummer, err1 = strconv.Atoi(lits[0])
	if len(lits[1]) > 0 {
		el.Wettkampfart = []rune(lits[1])[0]
	} else {
		err2 = fmt.Errorf("wettkampfart kein Zeichen")
	}
	el.Wertungsklasse = types.NewWertungsklasse(lits[2])
	el.MindestJahrgang = lits[3]
	el.MaximalJahrgang = lits[4]
	el.Pflichtzeit, err3 = types.NewZeit(lits[5])
	if len(lits[6]) > 0 {
		el.Geschlecht = types.NewGeschlecht([]rune(lits[6])[0])
	}

	return el, errors.Join(err1, err2, err3)
}
