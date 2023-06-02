package types

type Wertungsklasse string

const (
	JAHRGANG     = "JG"
	ALTERSKLASSE = "AK"
)

func NewWertungsklasse(value string) Wertungsklasse {
	return Wertungsklasse(value)
}
