package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/konrad2002/dsvparser/model"
	"github.com/konrad2002/dsvparser/parser"
)

func main() {
	BeispielNutzung()
}

func BeispielNutzung() {
	dat, err := os.ReadFile("assets/sdm26.dsv7")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(dat)
	r := parser.NewReader(buf)
	res, err := r.Read()
	if err != nil {
		panic(err)
	}
	def := res.(*model.Wettkampfergebnisliste)
	fmt.Printf(def.Veranstaltung.Veranstaltungsort)

	for _, p := range def.StaffelPerson {
		fmt.Printf("\n%s %s", p.Name, p.StartnummerSchwimmerStaffel)
	}
}
