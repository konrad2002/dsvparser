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
	args := 6
	if len(lits) != args {
		return Abschnitt{}, fmt.Errorf("zu wenig Argumente für ABSCHNITT, %d statt %d", len(lits), args)
	}
	var el Abschnitt
	var err, err1, err2 error

	el.Abschnittsnummer, err = strconv.Atoi(lits[0])
	if err != nil {
		return Abschnitt{}, fmt.Errorf("abschnittsnummer war nicht vom Typ 'Zahl'")
	}

	el.Abschnittsdatum, err1 = types.NewDatum(lits[1])
	el.Einlass, _ = types.NewUhrzeit(lits[2])
	el.Kampfrichtersitzung, _ = types.NewUhrzeit(lits[3])
	el.Anfangszeit, err2 = types.NewUhrzeit(lits[4])
	el.RelativeAngabe = lits[5] == "J"

	err = errors.Join(err1, err2)

	return el, err
}
