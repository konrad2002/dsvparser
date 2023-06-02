package elements

type Wettkampf struct {
	Wettkampfnummer               int
	Wettkampfart                  rune
	Abschnittsnummer              int
	AnzahlStarter                 int
	Einzelstrecke                 int
	Technik                       rune
	Ausuebung                     string
	Geschlecht                    rune
	ZuordnungBestenliste          string
	Qualifikationswettkampfnummer int
	Qualifikationswettkampfart    rune
}
