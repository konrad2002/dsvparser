package main

import (
	"fmt"
	"github.com/konrad2002/dsvparser/model"
	"github.com/konrad2002/dsvparser/model/elements"
	"io"
)

type Reader struct {
	p      *Parser
	format elements.Format
}

func NewReader(r io.Reader) *Reader {
	return &Reader{p: NewParser(r)}
}

func (r *Reader) Read() (model.Liste, error) {
	// skip comments and read format
	f := false
	for f == false {
		el, err := r.p.Parse()
		if err != nil {
			return nil, err
		}

		// comment
		if el == nil {
			continue
		}

		switch ele := el.(type) {
		case elements.Format:
			r.format = ele
			f = true
		default:
			return nil, fmt.Errorf("datei beginnt nicht mit FORMAT Element")
		}
	}

	if r.format.Version != 6 && r.format.Version != 7 {
		return nil, fmt.Errorf("version der Datei wird nicht unterstützt")
	}

	var list model.Liste

	switch r.format.Listart {
	case "Wettkampfdefinitionsliste":
		list = &model.Wettkampfdefinitionsliste{}
	case "Wettkampfergebnisliste":
		list = &model.Wettkampfergebnisliste{}
	case "Vereinsmeldeliste":
		return nil, fmt.Errorf("nicht implementiert")
	case "Vereinsergebnisliste":
		return nil, fmt.Errorf("nicht implementiert")
	default:
		return nil, fmt.Errorf("ungültige Listart")
	}
	list.AddElement(r.format)
	return r.readRestOfFile(list)
}

func (r *Reader) readRestOfFile(list model.Liste) (res model.Liste, err error) {
	res = list
	for {
		el, err := r.p.Parse()
		if err != nil {
			fmt.Printf("error\n")
			return res, err
		}

		// comment
		if el == nil {
			fmt.Printf("comment\n")
			continue
		}

		switch el.(type) {
		case elements.Dateiende:
			return res, nil
		default:
			res.AddElement(el)
		}
	}
}
