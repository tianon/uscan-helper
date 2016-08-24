package main

import (
	"bytes"
	"html/template"
)

type HtmlTemplate struct {
	Title    string
	Versions map[string]TemplateVersion
}

type TemplateVersion struct {
	Url string
}

var tmpl = template.Must(template.New("html").Parse(htmlTemplate))

func renderTemplate(data HtmlTemplate) (string, error) {
	out := &bytes.Buffer{}
	err := tmpl.Execute(out, data)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

const htmlTemplate = `
<!doctype html>
<html lang="en">
	<head>
		<meta charset="utf-8">

		<title>{{ .Title }}</title>
	</head>

	<body>
		<ul>
			{{- range $k, $v := .Versions }}
			<li><a href="{{ $v.Url }}">{{ $k }}</a></li>
			{{- end }}
		</ul>
	</body>
</html>
`
