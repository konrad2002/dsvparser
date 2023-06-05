package main

import (
	"bytes"
	"fmt"
	"github.com/konrad2002/dsvparser/model"
	"github.com/konrad2002/dsvparser/model/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReader_Read_WithoutComments(t *testing.T) {
	input := "FORMAT:Wettkampfdefinitionsliste;7;\nDATEIENDE"
	buf := bytes.NewBufferString(input)
	r := NewReader(buf)
	l, _ := r.Read()
	assert.IsType(t, &model.Wettkampfdefinitionsliste{}, l)
}

func TestReader_Read_WithComments(t *testing.T) {
	input := "(* bla *)\n(**)\nFORMAT:Wettkampfdefinitionsliste;7;\nDATEIENDE"
	buf := bytes.NewBufferString(input)
	r := NewReader(buf)
	l, _ := r.Read()
	assert.IsType(t, &model.Wettkampfdefinitionsliste{}, l)
}

func TestReader_Read_NotStartingWithFormat(t *testing.T) {
	input := "(* bla *)\nABSCHNITT:1;14.05.2023;;;09:00;;\nFORMAT:Wettkampfdefinitionsliste;7;\nDATEIENDE"
	buf := bytes.NewBufferString(input)
	r := NewReader(buf)
	_, err := r.Read()
	assert.Error(t, fmt.Errorf("datei beginnt nicht mit FORMAT Element"), err)
}

func TestReader_Read_WrongVersion(t *testing.T) {
	input := "(* bla *)\nFORMAT:Wettkampfdefinitionsliste;5;\nDATEIENDE"
	buf := bytes.NewBufferString(input)
	r := NewReader(buf)
	_, err := r.Read()
	assert.Error(t, fmt.Errorf("version der Datei wird nicht unterstützt"), err)
}

func TestReader_Read_WrongListart(t *testing.T) {
	input := "(* bla *)\nFORMAT:Dateiblödsinn;7;\nDATEIENDE"
	buf := bytes.NewBufferString(input)
	r := NewReader(buf)
	_, err := r.Read()
	assert.Error(t, fmt.Errorf("ungültige Listart"), err)
}

func TestReader_Read_NotImplemented(t *testing.T) {
	input := "(* bla *)\nFORMAT:Vereinsmeldeliste;7;\nDATEIENDE"
	buf := bytes.NewBufferString(input)
	r := NewReader(buf)
	_, err := r.Read()
	assert.Error(t, fmt.Errorf("nicht implementiert"), err)
}

func TestReader_Read_Definitionsliste_Type(t *testing.T) {
	input := "(* bla *)\nFORMAT:Wettkampfdefinitionsliste;7;\nABSCHNITT:1;14.05.2023;;;09:00;;\nDATEIENDE"
	buf := bytes.NewBufferString(input)
	r := NewReader(buf)
	res, _ := r.Read()
	assert.IsType(t, &model.Wettkampfdefinitionsliste{}, res)
}

func TestReader_Read_Definitionsliste_Entries(t *testing.T) {
	input := "(* bla *)\nFORMAT:WettkampfdeFINItionsliste;7;\nABSCHNITT:1;14.05.2023;;;09:00;;\nDATEIENDE"
	buf := bytes.NewBufferString(input)
	r := NewReader(buf)
	res, err := r.Read()
	if err != nil {
		fmt.Printf(err.Error())
	}
	def := res.(*model.Wettkampfdefinitionsliste)
	assert.Equal(t, 7, def.Format.Version)
	dat, _ := types.NewDatum("14.05.2023")
	assert.Equal(t, dat, def.Abschnitte[0].Abschnittsdatum)
}

func TestReader_Read_Ergebnisliste_Type(t *testing.T) {
	input := "(* bla *)\nFORMAT:WETTKAMPFERGEBNISLISTE;7;\nABSCHNITT:1;14.05.2023;;;09:00;;\nDATEIENDE"
	buf := bytes.NewBufferString(input)
	r := NewReader(buf)
	res, _ := r.Read()
	assert.IsType(t, &model.Wettkampfergebnisliste{}, res)
}

func TestReader_Read_Ergebnisliste_Entries(t *testing.T) {
	input := "(* bla *)\nFORMAT:WETTKAMPFERGEBNISliste;7;\nABSCHNITT:1;14.05.2023;;;09:00;;\nDATEIENDE"
	buf := bytes.NewBufferString(input)
	r := NewReader(buf)
	res, err := r.Read()
	if err != nil {
		fmt.Printf(err.Error())
	}
	def := res.(*model.Wettkampfergebnisliste)
	assert.Equal(t, 7, def.Format.Version)
	dat, _ := types.NewDatum("14.05.2023")
	assert.Equal(t, dat, def.Abschnitte[0].Abschnittsdatum)
}
