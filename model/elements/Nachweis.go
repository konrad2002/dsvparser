package elements

import (
	"dsvparser/model/types"
)

type Nachweis struct {
	NachweisVon types.Datum
	NachweisBis types.Datum
	Bahnlaenge  string
}
