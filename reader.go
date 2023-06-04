package main

import (
	"fmt"
	"github.com/konrad2002/dsvparser/model"
	"github.com/konrad2002/dsvparser/model/elements"
	"io"
)

type Reader struct {
	p *Parser
}

func NewReader(r io.Reader) *Reader {
	return &Reader{p: NewParser(r)}
}

func (r *Reader) Read() (interface{}, error) {
	var format elements.Format
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
			format = ele
			f = true
		default:
			return nil, fmt.Errorf("datei beginnt nicht mit FORMAT Element")
		}
	}

	if format.Version != 6 && format.Version != 7 {
		return nil, fmt.Errorf("version der Datei wird nicht unterstützt")
	}

	switch format.Listart {
	case "Wettkampfdefinitionsliste":
		return r.readDefinitionsliste()
	case "Wettkampfergebnisliste":
		return r.readErgebnisliste()
	case "Vereinsmeldeliste":
		return nil, fmt.Errorf("nicht implementiert")
	case "Vereinsergebnisliste":
		return nil, fmt.Errorf("nicht implementiert")
	default:
		return nil, fmt.Errorf("ungültige Listart")
	}
}

func (r *Reader) readDefinitionsliste() (model.Wettkampfdefinitionsliste, error) {
	// TODO parsen fortsetzen und Liste vervollständigen
	return model.Wettkampfdefinitionsliste{}, nil
}

func (r *Reader) readErgebnisliste() (model.Wettkampfergebnisliste, error) {
	// TODO parsen fortsetzen und Liste vervollständigen
	return model.Wettkampfergebnisliste{}, nil
}
