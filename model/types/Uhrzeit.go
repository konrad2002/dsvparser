package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Uhrzeit struct {
	Stunde int
	Minute int
}

func (u *Uhrzeit) getTime() time.Time {
	t, _ := time.Parse("15:04", u.getString())
	return t
}

func (u *Uhrzeit) getString() string {
	return fmt.Sprintf("%02d", u.Stunde) + ":" + fmt.Sprintf("%02d", u.Minute)
}

func NewUhrzeit(format string) Uhrzeit {
	var uhrzeit Uhrzeit
	var i = strings.Index(format, ":")
	uhrzeit.Stunde, _ = strconv.Atoi(format[:i])
	uhrzeit.Minute, _ = strconv.Atoi(format[i+1:])
	return uhrzeit
}
