package model

import "github.com/konrad2002/dsvparser/model/elements"

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

func (w *Wettkampfergebnisliste) AddElement(e interface{}) {
	switch v := e.(type) {
	case elements.Format:
		w.Format = v
	case elements.Erzeuger:
		w.Erzeuger = v
	case elements.Veranstaltung:
		w.Veranstaltung = v
	case elements.Veranstalter:
		w.Veranstalter = v
	case elements.Ausrichter:
		w.Ausrichter = v
	case elements.Abschnitt:
		w.Abschnitte = append(w.Abschnitte, v)
	case elements.Kampfgericht:
		w.Kampfgericht = append(w.Kampfgericht, v)
	case elements.Wettkampf:
		w.Wettkaempfe = append(w.Wettkaempfe, v)
	case elements.Wertung:
		w.Wertungen = append(w.Wertungen, v)
	case elements.Verein:
		w.Vereine = append(w.Vereine, v)
	case elements.PNErgebnis:
		w.PNErgebnisse = append(w.PNErgebnisse, v)
	case elements.PNZwischenzeit:
		w.PNZwischenzeiten = append(w.PNZwischenzeiten, v)
	case elements.PNReaktion:
		w.PNReaktionszeiten = append(w.PNReaktionszeiten, v)
	}
}
