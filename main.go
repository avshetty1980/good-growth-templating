package main

import (
	"log"

	"github.com/avshetty1980/good-growth-templating/lib"
)

type Templ map[string]string
type Arithmetic map[string]int

const (
	_ int = iota
	one
	two
	three
	four
	five
	six
	seven
	eight
	nine
)

func main() {

	templateString := "hello {{.foo | ToUpper}}\n"
	err := lib.TemplatePipeline(templateString, lib.Templ{"foo": "BaR"})

	if err != nil {
		log.Fatal("error executing template", err)
	}

	templateString = "hello {{.foo | ToUpper | ToLower}}\n"
	err = lib.TemplatePipeline(templateString, lib.Templ{"foo": "BAR"})

	if err != nil {
		log.Fatal("error executing template", err)
	}

	templateString = "{{.one | plus 2 3 4 5 }}\n"

	err = lib.TemplateArithmetic(templateString, lib.Arithmetic{"one": one})
	if err != nil {
		log.Fatal("error executing template", err)
	}

	templateString = "{{.one | mult 2 3 }}\n"
	err = lib.TemplateArithmetic(templateString, lib.Arithmetic{"one": one})
	if err != nil {
		log.Fatal("error executing template", err)
	}

}
