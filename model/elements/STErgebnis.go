package elements

import (
	"errors"
	"fmt"
	"github.com/konrad2002/dsvparser/model/types"
)

type STErgebnis struct {
	Wettkampfnummer                      types.Zahl
	Wettkampfart                         string
	WertungsID                           types.Zahl
	Platz                                types.Zahl
	GrundDerNichtwertung                 string
	NummerMannschaft                     types.Zahl
	VeranstaltungsIdStaffel              types.Zahl
	Verein                               string
	Vereinskennzahl                      types.Zahl
	Endzeit                              types.Zeit
	StartnummerDisqualifizierterSportler types.Zahl
	Disqualifikationsbemerkung           string
	ENM                                  string
}

func NewSTErgebnis(lits []string) (STErgebnis, error) {
	args := 13
	if len(lits) != args {
		return STErgebnis{}, fmt.Errorf("falsche Anzahl an Argumenten für STERGEBNIS, %d statt %d", len(lits), args)
	}
	var el STErgebnis
	var err1, err2, err3, err4, err5, err6, err7 error

	el.Wettkampfnummer, err1 = types.NewZahl(lits[0])
	el.Wettkampfart = lits[1]
	el.WertungsID, err2 = types.NewZahl(lits[2])
	el.Platz, err3 = types.NewZahl(lits[3])
	el.GrundDerNichtwertung = lits[4]
	el.NummerMannschaft, err4 = types.NewZahl(lits[5])
	el.VeranstaltungsIdStaffel, err4 = types.NewZahl(lits[6])
	el.Verein = lits[7]
	el.Vereinskennzahl, err5 = types.NewZahl(lits[8])
	el.Endzeit, err6 = types.NewZeit(lits[9])
	el.StartnummerDisqualifizierterSportler, err7 = types.NewZahl(lits[10])
	el.Disqualifikationsbemerkung = lits[11]
	el.ENM = lits[12]

	return el, errors.Join(err1, err2, err3, err4, err5, err6, err7)
}
