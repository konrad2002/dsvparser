package main

import (
	"bytes"
	"fmt"
	"github.com/konrad2002/dsvparser/model"
	"os"
)

func main() {
	fmt.Printf("nothing implemented for now...")
	exampleUsage()
}

type Example1 struct {
	Name string
}

type Example2 struct {
	Name string
}

func exampleUsage() {
	dat, err := os.ReadFile("assets/definition.dsv7")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(dat))

	var def model.Wettkampfdefinitionsliste

	buf := bytes.NewBuffer(dat)
	p := NewParser(buf)

	res, err := p.Parse()
	if err != nil {
		panic(err)
	}

	def.AddElement(res)

	var e Example2
	e.Name = "Paul"
	e2 := dodo(e)
	switch e2.(type) {
	case Example1:
		fmt.Printf("yep1")
		fmt.Printf(e2.(Example1).Name)
		break
	case Example2:
		fmt.Printf("yep2")
		fmt.Printf(e2.(Example2).Name)
		break

	}
}

func dodo(e Example2) interface{} {
	return e
}
