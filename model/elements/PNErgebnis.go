package elements

import "github.com/konrad2002/dsvparser/model/types"

type PNErgebnis struct {
	Wettkampfnummer            int
	Wettkampfart               string
	WertungsID                 int
	Platz                      int
	GrundDerNichtwertung       string
	Name                       string
	DsvId                      int
	VeranstaltungsIdSchwimmer  int
	Geschlecht                 types.Geschlecht
	Jahrgang                   int
	Altersklasse               int
	Verein                     string
	Vereinskennzahl            int
	Endzeit                    types.Zeit
	Disqualifikationsbemerkung string
	ENM                        rune
	Nationalitaet1             string
	Nationalitaet2             string
	Nationalitaet3             string
}

// TODO constructor
