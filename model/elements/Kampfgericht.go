package elements

import "dsvparser/model/types"

type Kampfgericht struct {
	Abschnittsnummer int
	Position         types.KampfrichterPosition
	Name             string
	Verein           string
}
