package types

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDatum(t *testing.T) {
	d, _ := NewDatum("16.04.2003")
	assert.Equal(t, 16, d.Tag)
	assert.Equal(t, 4, d.Monat)
	assert.Equal(t, 2003, d.Jahr)
}

func TestDatumGetTime(t *testing.T) {
	d, _ := NewDatum("16.04.2003")
	time := d.Time()
	assert.Equal(t, 16, time.Day())
	assert.Equal(t, "April", time.Month().String())
	assert.Equal(t, 2003, time.Year())
}

func TestDatumGetString(t *testing.T) {
	d, _ := NewDatum("16.04.2003")
	assert.Equal(t, "16.04.2003", d.String())
}

func TestNewDatumIllegalString(t *testing.T) {
	_, err := NewDatum("blaumeise")
	assert.Error(t, fmt.Errorf("datum ist nicht in der Form 'dd.dd.dddd'"), err)
}
