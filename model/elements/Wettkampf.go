package elements

import "github.com/konrad2002/dsvparser/model/types"

type Wettkampf struct {
	Wettkampfnummer               int
	Wettkampfart                  rune
	Abschnittsnummer              int
	AnzahlStarter                 int
	Einzelstrecke                 int
	Technik                       rune
	Ausuebung                     string
	Geschlecht                    types.Geschlecht
	ZuordnungBestenliste          string
	Qualifikationswettkampfnummer int
	Qualifikationswettkampfart    rune
}

// TODO: constructor
