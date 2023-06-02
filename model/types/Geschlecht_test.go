package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGeschlechtM(t *testing.T) {
	g := NewGeschlecht('M')
	assert.Equal(t, Geschlecht(MAENNLICH), g)
}

func TestNewGeschlechtW(t *testing.T) {
	g := NewGeschlecht('W')
	assert.Equal(t, Geschlecht(WEIBLICH), g)
}

func TestNewGeschlechtD(t *testing.T) {
	g := NewGeschlecht('D')
	assert.Equal(t, Geschlecht(DIVERS), g)
}
