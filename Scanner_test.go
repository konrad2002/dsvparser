package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScanner_Scan(t *testing.T) {
	input := "AUSRICHTER:Schwimmteam Erzgebirge;Steiner, Alexander;Hammergasse 24;09526;Olbernhau;GER;+491626802160;;alex@schwimmteamerzgebirge.de;\nABSCHNITT:1;14.05.2023;09:00;;"
	buf := bytes.NewBufferString(input)
	s := NewScanner(buf)
	tok, lit := s.Scan()
	assert.Equal(t, AUSRICHTER, tok)
	assert.Equal(t, "Hammergasse 24", lit[2])
	assert.Equal(t, 9, len(lit))
	tok, lit = s.Scan()
	assert.Equal(t, ABSCHNITT, tok)
	assert.Equal(t, 4, len(lit))
}

func TestScanner_ScanWithComment(t *testing.T) {
	input := "(*blub*)\nAUSRICHTER:Schwimmteam Erzgebirge;Steiner, Alexander;Hammergasse 24;09526;Olbernhau;GER;+491626802160;;alex@schwimmteamerzgebirge.de;\nABSCHNITT:1;14.05.2023;09:00;;"
	buf := bytes.NewBufferString(input)
	s := NewScanner(buf)
	tok, lit := s.Scan()
	assert.Equal(t, COMMENT, tok)
	tok, lit = s.Scan()
	assert.Equal(t, AUSRICHTER, tok)
	assert.Equal(t, "Hammergasse 24", lit[2])
	assert.Equal(t, 9, len(lit))
	tok, lit = s.Scan()
	assert.Equal(t, ABSCHNITT, tok)
	assert.Equal(t, 4, len(lit))
}
