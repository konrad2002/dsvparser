package elements

import "dsvparser/model/types"

type Pflichtzeit struct {
	Wettkampfnummer int
	Wettkampfart    rune
	Wertungsklasse  types.Wertungsklasse
	MindestJahrgang string
	MaximalJahrgang string
	Pflichtzeit     types.Zeit
	Geschlecht      types.Geschlecht
}
