package bench

import (
	"bytes"
	"html/template"
	"testing"
)

func BenchmarkTemplates(t *testing.B) {
	tpl := "Hello {{.Name}}"
	data := &map[string]string{
		"Name": "John Doe",
	}

	var buf bytes.Buffer
	for i := 0; i < t.N; i++ {
		tm, _ := template.New("test").Parse(tpl)
		tm.Execute(&buf, data)
		buf.Reset()
	}
}

func BenchmarkCompiledTemplates(t *testing.B) {
	tpl := "Hello {{.Name}}"
	tm, _ := template.New("test").Parse(tpl)
	data := &map[string]string{
		"Name": "John Doe",
	}

	var buf bytes.Buffer
	for i := 0; i < t.N; i++ {
		tm.Execute(&buf, data)
		buf.Reset()
	}
}

func BenchmarkParallelTemplates(t *testing.B) {
	tpl := "Hello {{.Name}}"
	tm, _ := template.New("test").Parse(tpl)
	data := &map[string]string{
		"Name": "John Doe",
	}

	t.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			tm.Execute(&buf, data)
			buf.Reset()
		}
	})
}
