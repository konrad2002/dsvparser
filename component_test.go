package main

import (
	"bytes"
	"fmt"
	"github.com/konrad2002/dsvparser/model"
	"github.com/konrad2002/dsvparser/model/types"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func Test_StandardExample_Definitionsliste(t *testing.T) {
	dat, err := os.ReadFile("assets/definition.dsv7")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(dat)
	r := NewReader(buf)
	res, _ := r.Read()
	def := res.(*model.Wettkampfdefinitionsliste)
	assert.Equal(t, 7, def.Format.Version)
	assert.Equal(t, "JAHRGANG 1990", def.Wertungen[2].Wertungsname)
	assert.Equal(t, strings.ToLower(types.EINZEL), strings.ToLower(def.Meldegelder[1].MeldegeldTyp))
}

func Test_StandardExample_Ergebnisliste(t *testing.T) {
	dat, err := os.ReadFile("assets/ergebnis.dsv7")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(dat)
	r := NewReader(buf)
	res, err := r.Read()
	if err != nil {
		fmt.Printf(err.Error())
	}
	erg := res.(*model.Wettkampfergebnisliste)
	assert.Equal(t, 7, erg.Format.Version)
	assert.Equal(t, "Duisburg", erg.Ausrichter.Ort)
	assert.Equal(t, 123440, erg.PNErgebnisse[4].DsvId)
}

func Test_BeispielMarienberg_Ergebnisliste(t *testing.T) {
	dat, err := os.ReadFile("assets/2023-05-14-Marienbe-Pr.DSV6")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(dat)
	r := NewReader(buf)
	res, err := r.Read()
	if err != nil {
		fmt.Printf(err.Error())
	}
	erg := res.(*model.Wettkampfergebnisliste)
	assert.Equal(t, 6, erg.Format.Version)
	assert.Equal(t, "Olbernhau", erg.Ausrichter.Ort)
	assert.Equal(t, 429663, erg.PNErgebnisse[4].DsvId)
}
