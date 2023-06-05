package elements

import (
	"errors"
	"fmt"
	"github.com/konrad2002/dsvparser/model/types"
	"strconv"
)

type Wettkampf struct {
	Wettkampfnummer               int
	Wettkampfart                  rune
	Abschnittsnummer              int
	AnzahlStarter                 int
	Einzelstrecke                 int
	Technik                       rune
	Ausuebung                     string
	Geschlecht                    types.Geschlecht
	ZuordnungBestenliste          string
	Qualifikationswettkampfnummer int
	Qualifikationswettkampfart    rune
}

func NewWettkampf(lits []string) (Wettkampf, error) {
	args := 11
	if len(lits) != args {
		return Wettkampf{}, fmt.Errorf("falsche Anzahl an Argumenten für WETTKAMPF, %d statt %d", len(lits), args)
	}
	var el Wettkampf
	var err1, err2, err3, err4, err5 error
	el.Wettkampfnummer, err1 = strconv.Atoi(lits[0])
	if len(lits[1]) == 1 {
		el.Wettkampfart = []rune(lits[1])[0]
	} else {
		err2 = fmt.Errorf("wettkampfart ist kein Zeichen")
	}
	el.Abschnittsnummer, err3 = strconv.Atoi(lits[2])
	el.AnzahlStarter, _ = strconv.Atoi(lits[3])
	el.Einzelstrecke, err4 = strconv.Atoi(lits[4])
	if len(lits[5]) == 1 {
		el.Technik = []rune(lits[5])[0]
	} else {
		err5 = fmt.Errorf("technik ist kein Zeichen")
	}
	el.Ausuebung = lits[6]
	if len(lits[7]) > 0 {
		el.Geschlecht = types.NewGeschlecht([]rune(lits[7])[0])
	}
	el.ZuordnungBestenliste = lits[8]
	el.Qualifikationswettkampfnummer, _ = strconv.Atoi(lits[9])
	if len(lits[10]) > 0 {
		el.Qualifikationswettkampfart = []rune(lits[10])[0]
	}
	return el, errors.Join(err1, err2, err3, err4, err5)
}
