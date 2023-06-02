package model

import "dsvparser/model/elements"

type Wettkampfergebnisliste struct {
	Format            elements.Format
	Erzeuger          elements.Erzeuger
	Veranstaltung     elements.Veranstaltung
	Veranstalter      elements.Veranstalter
	Ausrichter        elements.Ausrichter
	Abschnitte        []elements.Abschnitt
	Kampfgericht      []elements.Kampfgericht
	Wettkaempfe       []elements.Wettkampf
	Wertungen         []elements.Wertung
	Vereine           []elements.Verein
	PNErgebnisse      []elements.PNErgebnis
	PNZwischenzeiten  []elements.PNZwischenzeit
	PNReaktionszeiten []elements.PNReaktion
}
