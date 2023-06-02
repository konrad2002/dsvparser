package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMeldegeldTyp(t *testing.T) {
	m := NewMeldegeldTyp("Staffelmeldegeld")
	assert.Equal(t, MeldegeldTyp(STAFFEL), m)
}
