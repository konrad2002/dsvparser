package types

import (
	"strconv"
)

type Zahl int

func (z *Zahl) Int() int {
	return int(*z)
}

func NewZahl(value string) (Zahl, error) {
	if value == "" {
		return 0, nil
	}
	z, err := strconv.Atoi(value)
	return Zahl(z), err
}
