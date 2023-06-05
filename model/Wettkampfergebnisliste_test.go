package model

import (
	"github.com/konrad2002/dsvparser/model/elements"
	"github.com/konrad2002/dsvparser/model/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWettergebnisliste_AddElementSingle(t *testing.T) {
	var w = Wettkampfergebnisliste{}
	var e = elements.Veranstalter{Name: "Mario"}
	w.AddElement(e)
	assert.Equal(t, "Mario", w.Veranstalter.Name)
}

func TestWettkampfErgebnisliste_AddElementArray(t *testing.T) {
	var w = Wettkampfergebnisliste{}
	var e = elements.Wertung{
		Wettkampfnummer: 12,
		Wettkampfart:    1,
		WertungsID:      1,
		Wertungsklasse:  types.NewWertungsklasse("JG"),
		MindestJahrgang: "2002",
		MaximalJahrgang: "2002",
		Geschlecht:      types.NewGeschlecht('M'),
		Wertungsname:    "Jahrgang",
	}
	w.AddElement(e)
	w.AddElement(e)
	assert.Equal(t, "Jahrgang", w.Wertungen[0].Wertungsname)
	assert.Equal(t, "Jahrgang", w.Wertungen[1].Wertungsname)
}
