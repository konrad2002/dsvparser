package elements

import (
	"errors"
	"fmt"
	"github.com/konrad2002/dsvparser/model/types"
	"strconv"
)

type Abschnitt struct {
	Abschnittsnummer    int
	Abschnittsdatum     types.Datum
	Einlass             types.Uhrzeit
	Kampfrichtersitzung types.Uhrzeit
	Anfangszeit         types.Uhrzeit
	RelativeAngabe      bool
}

func NewAbschnitt(lits []string) (Abschnitt, error) {
	args1, args2 := 6, 4
	if len(lits) != args1 && len(lits) != args2 {
		return Abschnitt{}, fmt.Errorf("falsche Anzahl an Argumenten für ABSCHNITT, %d statt %d/%d", len(lits), args1, args2)
	}
	var el Abschnitt
	var err, err1, err2 error

	el.Abschnittsnummer, err = strconv.Atoi(lits[0])
	if err != nil {
		return Abschnitt{}, fmt.Errorf("abschnittsnummer war nicht vom Typ 'Zahl'")
	}

	el.Abschnittsdatum, err1 = types.NewDatum(lits[1])
	if len(lits) == 6 {
		el.Einlass, _ = types.NewUhrzeit(lits[2])
		el.Kampfrichtersitzung, _ = types.NewUhrzeit(lits[3])
	}
	el.Anfangszeit, err2 = types.NewUhrzeit(lits[len(lits)-2])
	el.RelativeAngabe = lits[len(lits)-1] == "J"

	err = errors.Join(err1, err2)

	return el, err
}
