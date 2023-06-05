package elements

import (
	"errors"
	"fmt"
	"github.com/konrad2002/dsvparser/model/types"
)

type Meldeschluss struct {
	Datum   types.Datum
	Uhrzeit types.Uhrzeit
}

func NewMeldeschluss(lits []string) (Meldeschluss, error) {
	args := 2
	if len(lits) != args {
		return Meldeschluss{}, fmt.Errorf("falsche Anzahl an Argumenten für MELDESCHLUSS, %d statt %d", len(lits), args)
	}
	var el Meldeschluss
	var err1, err2 error
	el.Datum, err1 = types.NewDatum(lits[0])
	el.Uhrzeit, err2 = types.NewUhrzeit(lits[1])

	return el, errors.Join(err1, err2)
}
