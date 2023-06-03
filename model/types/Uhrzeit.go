package types

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Uhrzeit struct {
	Stunde int
	Minute int
}

func (u *Uhrzeit) Time() time.Time {
	t, _ := time.Parse("15:04", u.String())
	return t
}

func (u *Uhrzeit) String() string {
	return fmt.Sprintf("%02d", u.Stunde) + ":" + fmt.Sprintf("%02d", u.Minute)
}

func NewUhrzeit(value string) (Uhrzeit, error) {
	ok, err := regexp.MatchString("\\d\\d:\\d\\d", value)
	if !ok || err != nil {
		return Uhrzeit{}, fmt.Errorf("uhrzeit ist nicht in der Form 'dd:dd'")
	}

	var uhrzeit Uhrzeit
	var i = strings.Index(value, ":")
	uhrzeit.Stunde, _ = strconv.Atoi(value[:i])
	uhrzeit.Minute, _ = strconv.Atoi(value[i+1:])
	return uhrzeit, nil
}
