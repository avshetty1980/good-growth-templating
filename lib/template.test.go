package lib

import (
	"testing"
)

func TestTemplatePipeline(t *testing.T) {
	got := TemplatePipeline("hello {{.foo | ToUpper | ToLower}}\n", Templ{"foo": "BAR"})

	if got != nil {
		t.Errorf("Error in TestTemplatePipeline %v", got)
	}
}

func TestTemplateArithmetic(t *testing.T) {
	got := TemplateArithmetic("{{.one | mult 2 3 }}\n", Arithmetic{"one": one})

	if got != nil {
		t.Errorf("Error in TestTemplatePipeline %v", got)
	}
}

func BenchmarkTemplPipeline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TemplatePipeline("hello {{.foo | ToUpper | ToLower}}\n", Templ{"foo": "BAR"})
	}
}

func BenchmarkTemplateArithmetic(t *testing.T) {
	got := TemplateArithmetic("{{.one | mult 2 3 }}\n", Arithmetic{"one": one})

	if got != nil {
		t.Errorf("Error in TestTemplatePipeline %v", got)
	}
}
