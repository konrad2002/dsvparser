package elements

import (
	"github.com/konrad2002/dsvparser/model/types"
)

type Abschnitt struct {
	Abschnittsnummer    int
	Abschnittsdatum     types.Datum
	Einlass             types.Uhrzeit
	Kampfrichtersitzung types.Uhrzeit
	Anfangszeit         types.Uhrzeit
	RelativeAngabe      bool
}
