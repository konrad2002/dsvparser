package parser

import (
	"bytes"
	"github.com/konrad2002/dsvparser/model/elements"
	"github.com/konrad2002/dsvparser/model/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	input := "AUSRICHTER:Schwimmteam Erzgebirge;Steiner, Alexander;Hammergasse 24;09526;Olbernhau;GER;+491626802160;;alex@schwimmteamerzgebirge.de;\nABSCHNITT:1;14.05.2023;09:00;;;;"
	buf := bytes.NewBufferString(input)
	p := NewParser(buf)
	el, _ := p.Parse()
	aus := el.(elements.Ausrichter)
	assert.Equal(t, "Hammergasse 24", aus.Strasse)
	el, _ = p.Parse()
	abs := el.(elements.Abschnitt)
	date, _ := types.NewDatum("14.05.2023")
	assert.Equal(t, 1, abs.Abschnittsnummer)
	assert.Equal(t, date, abs.Abschnittsdatum)
}
