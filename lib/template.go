package lib

import (
	"fmt"
	"os"
	"strings"
	"text/template"
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

var funcMap = template.FuncMap{
	"ToUpper": strings.ToUpper,
	"ToLower": strings.ToLower,
	"plus": func(i int, nums ...int) (int, error) {
		if len(nums) == 0 {
			return 0, fmt.Errorf("cannot have missing numbers")
		}
		total := i
		for _, num := range nums {

			total += num
		}
		return total, nil
	},
	"mult": func(i int, nums ...int) (int, error) {
		if len(nums) == 0 {
			return 0, fmt.Errorf("cannot have missing numbers")
		}
		total := i
		for _, num := range nums {
			total *= num
		}
		return total, nil
	},
}

func TemplatePipeline(templStr string, str Templ) error {
	t := template.Must(template.New("templ").Funcs(funcMap).Parse(templStr))
	err := t.Execute(os.Stdout, str)

	if err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}
	return nil
}

func TemplateArithmetic(templStr string, num Arithmetic) error {
	t := template.Must(template.New("templ").Funcs(funcMap).Parse(templStr))
	err := t.Execute(os.Stdout, num)

	if err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}
	return nil
}
