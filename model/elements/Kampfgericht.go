package elements

import "github.com/konrad2002/dsvparser/model/types"

type Kampfgericht struct {
	Abschnittsnummer int
	Position         types.KampfrichterPosition
	Name             string
	Verein           string
}

// TODO: constructor
