package main

import (
	"bytes"
	"fmt"
	"github.com/konrad2002/dsvparser/model"
	"github.com/konrad2002/dsvparser/parser"
	"os"
)

func main() {
	BeispielNutzung()
}

func BeispielNutzung() {
	dat, err := os.ReadFile("assets/definition.dsv7")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(dat)
	r := parser.NewReader(buf)
	res, err := r.Read()
	if err != nil {
		panic(err)
	}
	def := res.(*model.Wettkampfdefinitionsliste)
	fmt.Printf(def.Veranstaltungsort.PLZ)
}
