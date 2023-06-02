package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDatum(t *testing.T) {
	d := NewDatum("16.04.2003")
	assert.Equal(t, 16, d.Tag)
	assert.Equal(t, 4, d.Monat)
	assert.Equal(t, 2003, d.Jahr)
}

func TestDatumGetTime(t *testing.T) {
	d := NewDatum("16.04.2003")
	time := d.GetTime()
	assert.Equal(t, 16, time.Day())
	assert.Equal(t, "April", time.Month().String())
	assert.Equal(t, 2003, time.Year())
}

func TestDatumGetString(t *testing.T) {
	d := NewDatum("16.04.2003")
	assert.Equal(t, "16.04.2003", d.GetString())
}
