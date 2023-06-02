package elements

import "dsvparser/model/types"

type PNReaktion struct {
	VeranstaltungsIdSchwimmer int
	Wettkampfnummer           int
	Wettkampfart              string
	Art                       rune
	Reaktionszeit             types.Zeit
}
