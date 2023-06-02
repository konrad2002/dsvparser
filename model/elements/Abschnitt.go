package elements

import (
	"dsvparser/model/types"
)

type Abschnitt struct {
	Abschnittsnummer    int
	Abschnittsdatum     types.Datum
	Einlass             types.Uhrzeit
	Kampfrichtersitzung types.Uhrzeit
	Anfangszeit         types.Uhrzeit
	RelativeAngabe      bool
}
