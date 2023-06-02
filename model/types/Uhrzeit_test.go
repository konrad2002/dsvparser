package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUhrzeit(t *testing.T) {
	u := NewUhrzeit("14:23")
	assert.Equal(t, 14, u.Stunde)
	assert.Equal(t, 23, u.Minute)
}

func TestNewUhrzeitWithLeadingZero(t *testing.T) {
	u := NewUhrzeit("01:02")
	assert.Equal(t, 1, u.Stunde)
	assert.Equal(t, 2, u.Minute)
}

func TestUhrzeitGetTime(t *testing.T) {
	u := NewUhrzeit("14:23")
	time := u.getTime()
	assert.Equal(t, 14, time.Hour())
	assert.Equal(t, 23, time.Minute())
}

func TestUhrzeitGetString(t *testing.T) {
	u := NewUhrzeit("15:13")
	assert.Equal(t, "15:13", u.getString())
}

func TestUhrzeitGetStringWithLeadingZero(t *testing.T) {
	u := NewUhrzeit("03:01")
	assert.Equal(t, "03:01", u.getString())
}
