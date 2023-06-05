package main

import (
	"bytes"
	"fmt"
	"github.com/konrad2002/dsvparser/model"
	"os"
)

func main() {
	exampleUsage()
}

func exampleUsage() {
	dat, err := os.ReadFile("assets/definition.dsv7")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(dat)
	r := NewReader(buf)
	res, err := r.Read()
	if err != nil {
		panic(err)
	}
	def := res.(*model.Wettkampfdefinitionsliste)
	fmt.Printf("%+v\n", def)
}
