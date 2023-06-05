package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTokenValid(t *testing.T) {
	tok := NewToken("DATEIENDE")
	assert.Equal(t, DATEIENDE, tok)
}

func TestNewTokenIllegal(t *testing.T) {
	tok := NewToken("blub")
	assert.Equal(t, ILLEGAL, tok)
}
