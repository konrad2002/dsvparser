package elements

import "dsvparser/model/types"

type PNZwischenzeit struct {
	VeranstaltungsIdSchwimmer int
	Wettkampfnummer           int
	Wettkampfart              string
	Distanz                   int
	Zwischenzeit              types.Zeit
}
