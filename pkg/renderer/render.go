package opencanvasRender

import (
	_ "embed"
	"html/template"
	"io"

	"github.com/irth/opencanvas-go/pkg/opencanvas"
)

//go:embed template.html
var templateString string
var tmpl *template.Template = template.Must(template.New("").Parse(templateString))

func Render(w io.Writer, canvas *opencanvas.Canvas) error {
	return tmpl.Execute(w, canvas)
}
