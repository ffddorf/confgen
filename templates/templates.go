package templates

import (
	_ "embed"
	"errors"
	"io"
	"text/template"

	"github.com/ffddorf/confgen/netbox"
)

var (
	// all known templates by name
	templates = map[string]*template.Template{}

	// common template functions
	funcs = template.FuncMap{
		"maybeQuote":              maybeQuote,
		"edgeosConfigFromMap":     edgeosConfigFromMap,
		"edgeosPrepareInterfaces": edgeosPrepareInterfaces,
	}
)

// template bodies embedded via `go:embed`
var (
	//go:embed edgeos.tpl
	edgeosRaw string
	// ... load your template file here!
)

func init() {
	initTemplate("edgeos", edgeosRaw)
	// ... initialize your template instance here!
}

// used to initialize a template globally using its name and body
func initTemplate(name string, tpl string) {
	templates[name] = template.Must(
		template.New(name).
			Funcs(funcs).
			Parse(tpl),
	)
}

// TemplateData is the data needed to execute a template
type TemplateData struct {
	Device *netbox.Device
}

var ErrTemplateNotFound = errors.New("template not found")

// Render executes the template by the given name and
// writes the result into `out`.
func Render(out io.Writer, name string, data TemplateData) error {
	templ, ok := templates[name]
	if !ok {
		return ErrTemplateNotFound
	}
	return templ.Execute(out, data)
}
