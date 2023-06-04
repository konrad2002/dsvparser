package elements

import (
	"fmt"
	"strconv"
)

type Format struct {
	Listart string
	Version int
}

func NewFormat(lits []string) (Format, error) {
	args := 2
	if len(lits) != args {
		return Format{}, fmt.Errorf("ungütlige Attribute für FORMAT, %d statt %d", len(lits), args)
	}
	v, err := strconv.Atoi(lits[1])
	if err != nil {
		return Format{}, err
	}
	return Format{Listart: lits[0], Version: v}, nil
}
