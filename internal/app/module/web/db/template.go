package db

import "text/template"

var (
	modelTemplate *template.Template
)

func init() {
	modelTemplate, _ = template.New("modelTemplate").Parse(modelTpl)
}

func SetModelTemplate(tmpl *template.Template) {
	if tmpl != nil {
		modelTemplate = tmpl
	}
}

var modelTpl = `
package {{ .PackageName }}

type {{ .ModelName }} struct {
	{{ .Fields }}
}

{{ if .Option.WithGormAnnotation }}
	{{ template "tableFunc" . }}
{{ end }}

{{ if .Option.WithGormAnnotation }}
 	{{ template "curdFunc" . }}
{{ end }}


`
