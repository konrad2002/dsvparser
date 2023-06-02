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

func (w *Wettkampfdefinitionsliste) AddElement(e interface{}) {
	switch v := e.(type) {
	case elements.Format:
		w.Format = v
	case elements.Erzeuger:
		w.Erzeuger = v
	case elements.Veranstaltung:
		w.Veranstaltung = v
	case elements.Veranstaltungsort:
		w.Veranstaltungsort = v
	case elements.AusschreibungImNetz:
		w.AusschreibungImNetz = v
	case elements.Veranstalter:
		w.Veranstalter = v
	case elements.Ausrichter:
		w.Ausrichter = v
	case elements.Meldeadresse:
		w.Meldeadresse = v
	case elements.Meldeschluss:
		w.Meldeschluss = v
	case elements.Bankverbindung:
		w.Bankverbindung = v
	case elements.Besonderes:
		w.Besonderes = v
	case elements.Nachweis:
		w.Nachweis = v
	case elements.Abschnitt:
		w.Abschnitte = append(w.Abschnitte, v)
	case elements.Wettkampf:
		w.Wettkaempfe = append(w.Wettkaempfe, v)
	case elements.Wertung:
		w.Wertungen = append(w.Wertungen, v)
	case elements.Pflichtzeit:
		w.Pflichtzeiten = append(w.Pflichtzeiten, v)
	case elements.Meldegeld:
		w.Meldegelder = append(w.Meldegelder, v)
	}
}
