package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Datum struct {
	Tag   int
	Monat int
	Jahr  int
}

func (d *Datum) GetTime() time.Time {
	t, _ := time.Parse("02.01.2006", d.GetString())
	return t
}

func (d *Datum) GetString() string {
	return fmt.Sprintf("%02d", d.Tag) + "." + fmt.Sprintf("%02d", d.Monat) + "." + fmt.Sprintf("%04d", d.Jahr)
}

func NewDatum(value string) Datum {
	var datum Datum
	var i = strings.Index(value, ".")
	datum.Tag, _ = strconv.Atoi(value[:i])
	value = value[i+1:]
	i = strings.Index(value, ".")
	datum.Monat, _ = strconv.Atoi(value[:i])
	datum.Jahr, _ = strconv.Atoi(value[i+1:])
	return datum
}
