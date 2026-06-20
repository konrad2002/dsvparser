package elements

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/konrad2002/dsvparser/model/types"
)

type StaffelPerson struct {
	VeranstaltungsIdStaffel     int
	Wettkampfnummer             int
	Wettkampfart                string
	Name                        string
	DsvId                       int
	StartnummerSchwimmerStaffel int
	Geschlecht                  types.Geschlecht
	Jahrgang                    int
	Altersklasse                int
	Nationalitaet1              string
	Nationalitaet2              string
	Nationalitaet3              string
}

func NewStaffelPerson(lits []string) (StaffelPerson, error) {
	args6 := 9
	args7 := 12
	if len(lits) != args6 && len(lits) != args7 {
		return StaffelPerson{}, fmt.Errorf("falsche Anzahl an Argumenten für STAFFELPERSON, %d statt %d/%d", len(lits), args6, args7)
	}
	var el StaffelPerson
	var err1, err2, err3, err4, err5, err6 error

	el.VeranstaltungsIdStaffel, err1 = strconv.Atoi(lits[0])
	el.Wettkampfnummer, err2 = strconv.Atoi(lits[1])
	el.Wettkampfart = lits[2]
	el.Name = lits[3]
	el.DsvId, err3 = strconv.Atoi(lits[4])
	el.StartnummerSchwimmerStaffel, err4 = strconv.Atoi(lits[5])
	if len(lits[6]) > 0 {
		el.Geschlecht = types.NewGeschlecht([]rune(lits[6])[0])
	} else {
		err5 = fmt.Errorf("geschlecht ist kein Zeichen")
	}
	el.Jahrgang, err6 = strconv.Atoi(lits[7])
	el.Altersklasse, _ = strconv.Atoi(lits[8])

	// nur DSV7
	if len(lits) == 12 {
		el.Nationalitaet1 = lits[9]
		el.Nationalitaet2 = lits[10]
		el.Nationalitaet3 = lits[11]
	}

	return el, errors.Join(err1, err2, err3, err4, err5, err6)
}
