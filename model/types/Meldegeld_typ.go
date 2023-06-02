package types

type MeldegeldTyp string

const (
	PAUSCHALE  = "Meldegeldpauschale"
	EINZEL     = "Einzelmeldegeld"
	STAFFEL    = "Staffelmeldegeld"
	WETTKAMPF  = "Wkmeldegeld"
	MANNSCHAFT = "Mannschaftmeldegeld"
)

func NewMeldegeldTyp(value string) MeldegeldTyp {
	return MeldegeldTyp(value)
}
