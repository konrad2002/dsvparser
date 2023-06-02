package model

import "github.com/konrad2002/dsvparser/model/elements"

type Wettkampfdefinitionsliste struct {
	Format              elements.Format
	Erzeuger            elements.Erzeuger
	Veranstaltung       elements.Veranstaltung
	Veranstaltungsort   elements.Veranstaltungsort
	AusschreibungImNetz elements.AusschreibungImNetz
	Veranstalter        elements.Veranstalter
	Ausrichter          elements.Ausrichter
	Meldeadresse        elements.Meldeadresse
	Meldeschluss        elements.Meldeschluss
	Bankverbindung      elements.Bankverbindung
	Besonderes          elements.Besonderes
	Nachweis            elements.Nachweis
	Abschnitte          []elements.Abschnitt
	Wettkaempfe         []elements.Wettkampf
	Wertungen           []elements.Wertung
	Pflichtzeiten       []elements.Pflichtzeit
	Meldegelder         []elements.Meldegeld
}
