package elements

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/konrad2002/dsvparser/model/types"
)

type STErgebnis struct {
	Wettkampfnummer            int
	Wettkampfart               string
	WertungsID                 int
	Platz                      int
	GrundDerNichtwertung       string
	NummerDerMannschaft        int
	VeranstaltungsIdStaffel    int
	Verein                     string
	Vereinskennzahl            int
	Endzeit                    types.Zeit
	StartnummerDisqSchwimmer   int
	Disqualifikationsbemerkung string
	ENM                        rune
}

func NewSTErgebnis(lits []string) (STErgebnis, error) {
	args6 := 10
	args7 := 13
	if len(lits) != args6 && len(lits) != args7 {
		return STErgebnis{}, fmt.Errorf("falsche Anzahl an Argumenten für STERGEBNIS, %d statt %d/%d", len(lits), args6, args7)
	}
	var el STErgebnis
	var err1, err2, err3, err4, err5, err6, err7 error

	el.Wettkampfnummer, err1 = strconv.Atoi(lits[0])
	el.Wettkampfart = lits[1]
	el.WertungsID, err2 = strconv.Atoi(lits[2])
	el.Platz, err3 = strconv.Atoi(lits[3])
	el.GrundDerNichtwertung = lits[4]
	el.NummerDerMannschaft, err4 = strconv.Atoi(lits[5])
	el.VeranstaltungsIdStaffel, err5 = strconv.Atoi(lits[6])
	el.Verein = lits[7]
	el.Vereinskennzahl, err6 = strconv.Atoi(lits[8])
	el.Endzeit, err7 = types.NewZeit(lits[9])

	// nur DSV7
	if len(lits) == 13 {
		el.StartnummerDisqSchwimmer, _ = strconv.Atoi(lits[10])
		el.Disqualifikationsbemerkung = lits[11]
		if len(lits[12]) > 0 {
			el.ENM = []rune(lits[12])[0]
		}
	}

	return el, errors.Join(err1, err2, err3, err4, err5, err6, err7)
}
