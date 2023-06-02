package elements

import "github.com/konrad2002/dsvparser/model/types"

type Wertung struct {
	Wettkampfnummer int
	Wettkampfart    rune
	WertungsID      int
	Wertungsklasse  types.Wertungsklasse
	MindestJahrgang string
	MaximalJahrgang string
	Geschlecht      types.Geschlecht
	Wertungsname    string
}
