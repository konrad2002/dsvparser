package types

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Zeit struct {
	Stunde      int
	Minute      int
	Sekunde     int
	Hundertstel int
}

func (z *Zeit) Duration() time.Duration {
	d, _ := time.ParseDuration(
		strconv.Itoa(z.Stunde) + "h" +
			strconv.Itoa(z.Minute) + "m" +
			strconv.Itoa(z.Sekunde) + "s" +
			fmt.Sprintf("%03d", z.Hundertstel*10) + "ms")
	return d
}

func (z *Zeit) String() string {
	return fmt.Sprintf("%02d", z.Stunde) + ":" +
		fmt.Sprintf("%02d", z.Minute) + ":" +
		fmt.Sprintf("%02d", z.Sekunde) + "," +
		fmt.Sprintf("%02d", z.Hundertstel)
}

func NewZeit(value string) (Zeit, error) {
	ok, err := regexp.MatchString("\\d\\d:\\d\\d:\\d\\d,\\d\\d", value)
	if !ok || err != nil {
		return Zeit{}, fmt.Errorf("zeit ist nicht in der Form 'dd:dd:dd,dd'")
	}

	var zeit Zeit
	var i = strings.Index(value, ":")
	zeit.Stunde, _ = strconv.Atoi(value[:i])
	value = value[i+1:]
	i = strings.Index(value, ":")
	zeit.Minute, _ = strconv.Atoi(value[:i])
	value = value[i+1:]
	i = strings.Index(value, ",")
	zeit.Sekunde, _ = strconv.Atoi(value[:i])
	zeit.Hundertstel, _ = strconv.Atoi(value[i+1:])
	return zeit, nil
}
