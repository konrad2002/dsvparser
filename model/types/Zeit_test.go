package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewZeit(t *testing.T) {
	z := NewZeit("01:15:54,87")
	assert.Equal(t, 1, z.Stunde)
	assert.Equal(t, 15, z.Minute)
	assert.Equal(t, 54, z.Sekunde)
	assert.Equal(t, 87, z.Hundertstel)
}

func TestZeitGetString(t *testing.T) {
	z := NewZeit("01:15:54,87")
	assert.Equal(t, "01:15:54,87", z.getString())
}

func TestZeitGetDuration(t *testing.T) {
	z := NewZeit("01:15:54,87")
	assert.InDelta(t, 1.265241, z.getDuration().Hours(), 0.001)
	assert.InDelta(t, 75.91, z.getDuration().Minutes(), 0.1)
	assert.InDelta(t, 4554.87, z.getDuration().Seconds(), 0.1)
	assert.Equal(t, int64(4554870), z.getDuration().Milliseconds())
}
