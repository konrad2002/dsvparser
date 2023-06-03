package elements

import "github.com/konrad2002/dsvparser/model/types"

type PNReaktion struct {
	VeranstaltungsIdSchwimmer int
	Wettkampfnummer           int
	Wettkampfart              string
	Art                       rune
	Reaktionszeit             types.Zeit
}

// TODO constructor
