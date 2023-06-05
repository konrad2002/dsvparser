package types

type KampfrichterPosition string

const (
	SCHIEDSRICHTER          = "SCH"
	STARTER                 = "STA"
	ZIELRICHTEROBMANN       = "ZRO"
	ZIELRICHTER             = "ZR"
	ZEITNEHMEROBMANN        = "ZNO"
	ZEITNEHMER              = "ZN"
	RESERVEZEITNEHMER       = "RZN"
	SCHWIMMRICHTER          = "SR"
	WENDERICHTEROBMANN      = "WRO"
	WENDERICHTER            = "WR"
	AUSWERTER               = "AUS"
	SPRECHER                = "SP"
	PROTOKOLLFUEHRER        = "PKF"
	STARTORDNER             = "STO"
	WETTKAMPFHELFER         = "WKH"
	ASSISTENZSCHIEDSRICHTER = "ASCH"
	SICHERHEITSBEAUFTRAGTER = "SIB"
	STRECKENAUFSICHT        = "SAUF"
	ORDNERVERSORGUNGSSTELLE = "VER"
	SONSTIGE                = "ZBV"
)

func NewKampfrichterPosition(value string) KampfrichterPosition {
	return KampfrichterPosition(value)
}
