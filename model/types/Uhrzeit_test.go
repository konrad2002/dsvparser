package types

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUhrzeit(t *testing.T) {
	u, _ := NewUhrzeit("14:23")
	assert.Equal(t, 14, u.Stunde)
	assert.Equal(t, 23, u.Minute)
}

func TestNewUhrzeitWithLeadingZero(t *testing.T) {
	u, _ := NewUhrzeit("01:02")
	assert.Equal(t, 1, u.Stunde)
	assert.Equal(t, 2, u.Minute)
}

func TestUhrzeitGetTime(t *testing.T) {
	u, _ := NewUhrzeit("14:23")
	time := u.Time()
	assert.Equal(t, 14, time.Hour())
	assert.Equal(t, 23, time.Minute())
}

func TestUhrzeitGetString(t *testing.T) {
	u, _ := NewUhrzeit("15:13")
	assert.Equal(t, "15:13", u.String())
}

func TestUhrzeitGetStringWithLeadingZero(t *testing.T) {
	u, _ := NewUhrzeit("03:01")
	assert.Equal(t, "03:01", u.String())
}

func TestNewUhrzeitIllegal(t *testing.T) {
	_, err := NewUhrzeit("blaumeise")
	assert.Error(t, fmt.Errorf("uhrzeit ist nicht in der Form 'dd:dd'"), err)
}
