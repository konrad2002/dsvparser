package types

type Geschlecht rune

const (
	MAENNLICH = 'M'
	WEIBLICH  = 'W'
	DIVERS    = 'D'
)

func NewGeschlecht(value rune) Geschlecht {
	return Geschlecht(value)
}
