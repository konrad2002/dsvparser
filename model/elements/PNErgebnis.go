package elements

import (
	"errors"
	"fmt"
	"github.com/konrad2002/dsvparser/model/types"
	"strconv"
)

type PNErgebnis struct {
	Wettkampfnummer            int
	Wettkampfart               string
	WertungsID                 int
	Platz                      int
	GrundDerNichtwertung       string
	Name                       string
	DsvId                      string
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

func NewPNErgebnis(lits []string) (PNErgebnis, error) {
	args6 := 16
	args7 := 19
	if len(lits) != args6 && len(lits) != args7 {
		return PNErgebnis{}, fmt.Errorf("falsche Anzahl an Argumenten für PNERGEBNIS, %d statt %d/%d", len(lits), args6, args7)
	}
	var el PNErgebnis
	var err1, err2, err3, err4, err5, err6, err7, err8, err9 error

	el.Wettkampfnummer, err1 = strconv.Atoi(lits[0])
	el.Wettkampfart = lits[1]
	el.WertungsID, err2 = strconv.Atoi(lits[2])
	el.Platz, err3 = strconv.Atoi(lits[3])
	el.GrundDerNichtwertung = lits[4]
	el.Name = lits[5]
	el.DsvId = lits[6]
	el.VeranstaltungsIdSchwimmer, err5 = strconv.Atoi(lits[7])
	if len(lits[8]) > 0 {
		el.Geschlecht = types.NewGeschlecht([]rune(lits[8])[0])
	} else {
		err6 = fmt.Errorf("geschlecht ist kein Zeichen")
	}
	el.Jahrgang, err7 = strconv.Atoi(lits[9])
	el.Altersklasse, _ = strconv.Atoi(lits[10])
	el.Verein = lits[11]
	el.Vereinskennzahl, err8 = strconv.Atoi(lits[12])
	el.Endzeit, err9 = types.NewZeit(lits[13])
	el.Disqualifikationsbemerkung = lits[14]
	if len(lits[15]) > 0 {
		el.ENM = []rune(lits[15])[0]
	}

	// nur DSV7
	if len(lits) == 19 {
		el.Nationalitaet1 = lits[16]
		el.Nationalitaet2 = lits[17]
		el.Nationalitaet3 = lits[18]
	}

	return el, errors.Join(err1, err2, err3, err4, err5, err6, err7, err8, err9)
}
