package elements

import (
	"fmt"
	"github.com/konrad2002/dsvparser/model/types"
)

type Nachweis struct {
	NachweisVon types.Datum
	NachweisBis types.Datum
	Bahnlaenge  string
}

func NewNachweis(lits []string) (Nachweis, error) {
	args := 3
	if len(lits) != args {
		return Nachweis{}, fmt.Errorf("zu wenig Argumente für NACHWEIS, %d statt %d", len(lits), args)
	}
	var el Nachweis
	var err error
	el.NachweisVon, err = types.NewDatum(lits[0])
	el.NachweisBis, _ = types.NewDatum(lits[1])
	el.Bahnlaenge = lits[2]

	return el, err
}
