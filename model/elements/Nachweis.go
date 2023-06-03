package elements

import (
	"github.com/konrad2002/dsvparser/model/types"
)

type Nachweis struct {
	NachweisVon types.Datum
	NachweisBis types.Datum
	Bahnlaenge  string
}

// TODO constructor
