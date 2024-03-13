package opencanvasRender

import (
	_ "embed"
	"html/template"
	"io"

	"github.com/irth/opencanvas-go/pkg/opencanvas"
)

//go:embed template.html
var templateString string

//go:embed normalize.css
var normalizeCSS string

func Normalize() template.CSS {
	return template.CSS(normalizeCSS)
}

var funcMap template.FuncMap = template.FuncMap{
	"normalize": Normalize,
}
var tmpl *template.Template = template.Must(template.New("").Funcs(funcMap).Parse(templateString))

func Render(w io.Writer, canvas *opencanvas.Canvas) error {
	return tmpl.Execute(w, canvas)
}
