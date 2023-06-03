package elements

import "github.com/konrad2002/dsvparser/model/types"

type PNZwischenzeit struct {
	VeranstaltungsIdSchwimmer int
	Wettkampfnummer           int
	Wettkampfart              string
	Distanz                   int
	Zwischenzeit              types.Zeit
}

// TODO constructor
