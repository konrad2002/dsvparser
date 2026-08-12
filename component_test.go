package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/konrad2002/dsvparser/model"
	"github.com/konrad2002/dsvparser/model/types"
	"github.com/konrad2002/dsvparser/parser"
	"github.com/stretchr/testify/assert"
)

func Test_StandardExample_Definitionsliste(t *testing.T) {
	dat, err := os.ReadFile("assets/definition.dsv7")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(dat)
	r := parser.NewReader(buf)
	res, _ := r.Read()
	def := res.(*model.Wettkampfdefinitionsliste)
	assert.Equal(t, 7, def.Format.Version)
	assert.Equal(t, false, def.Lastschrift.Hinweis)
	assert.Equal(t, "JAHRGANG 1990", def.Wertungen[2].Wertungsname)
	assert.Equal(t, strings.ToLower(types.EINZEL), strings.ToLower(def.Meldegelder[1].MeldegeldTyp))
}

func Test_StandardExample_Definitionsliste_Version_8(t *testing.T) {
	dat, err := os.ReadFile("assets/definition.dsv8")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(dat)
	r := parser.NewReader(buf)
	res, _ := r.Read()
	def := res.(*model.Wettkampfdefinitionsliste)
	assert.Equal(t, 8, def.Format.Version)
	assert.Equal(t, true, def.Lastschrift.Hinweis)
	assert.Equal(t, "KB", def.Wettkaempfe[0].Ausuebung)
	assert.Equal(t, "KR", def.Wettkaempfe[2].Ausuebung)
	assert.Equal(t, strings.ToLower(types.TEILNEHMER), strings.ToLower(def.Meldegelder[1].MeldegeldTyp))
	assert.Equal(t, strings.ToLower(types.ABSCHNITT), strings.ToLower(def.Meldegelder[4].MeldegeldTyp))
}

func Test_StandardExample_Definitionsliste_Version_9(t *testing.T) {
	dat, err := os.ReadFile("assets/definition.dsv9")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(dat)
	r := parser.NewReader(buf)
	_, err2 := r.Read()
	assert.Equal(t, "version der Datei wird nicht unterstützt", err2.Error())
}

func Test_StandardExample_Ergebnisliste(t *testing.T) {
	dat, err := os.ReadFile("assets/ergebnis.dsv7")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(dat)
	r := parser.NewReader(buf)
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
	r := parser.NewReader(buf)
	res, err := r.Read()
	if err != nil {
		fmt.Printf(err.Error())
	}
	erg := res.(*model.Wettkampfergebnisliste)
	assert.Equal(t, 6, erg.Format.Version)
	assert.Equal(t, "Olbernhau", erg.Ausrichter.Ort)
	assert.Equal(t, 429663, erg.PNErgebnisse[4].DsvId)
}

func Test_BeispielIESC_Ergebnisliste(t *testing.T) {
	dat, err := os.ReadFile("assets/Ergebnisdatei.dsv6")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(dat)
	r := parser.NewReader(buf)
	res, err := r.Read()
	if err != nil {
		fmt.Printf(err.Error())
	}
	erg := res.(*model.Wettkampfergebnisliste)
	assert.Equal(t, 6, erg.Format.Version)
	assert.Equal(t, "", erg.Ausrichter.Ort)
}
