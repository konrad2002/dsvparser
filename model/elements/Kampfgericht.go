package elements

import (
	"errors"
	"fmt"
	"github.com/konrad2002/dsvparser/model/types"
	"strconv"
)

type Kampfgericht struct {
	Abschnittsnummer int
	Position         types.KampfrichterPosition
	Name             string
	Verein           string
}

func NewKampfgericht(lits []string) (Kampfgericht, error) {
	args := 4
	if len(lits) != args {
		return Kampfgericht{}, fmt.Errorf("zu wenig Argumente für KAMPFGERICHT, %d statt %d", len(lits), args)
	}
	var el Kampfgericht
	var err, err1, err2 error

	el.Abschnittsnummer, err = strconv.Atoi(lits[0])
	el.Position = types.NewKampfrichterPosition(lits[1])
	el.Name = lits[2]
	el.Verein = lits[3]

	err = errors.Join(err1, err2)
	return el, err
}
