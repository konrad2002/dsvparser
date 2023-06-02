package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWertungsklasseJG(t *testing.T) {
	w := NewWertungsklasse("JG")
	assert.Equal(t, Wertungsklasse(JAHRGANG), w)
}

func TestNewWertungsklasseAK(t *testing.T) {
	w := NewWertungsklasse("AK")
	assert.Equal(t, Wertungsklasse(ALTERSKLASSE), w)
}
