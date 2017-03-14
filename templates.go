package main

import "html/template"

var indexTemplate *template.Template = nil

func GetIndexTemplate() *template.Template {
	if FileportConfig.RecompileTemplates || indexTemplate == nil {
		indexTemplate, _ = template.ParseFiles("templates/index.html")
	}

	return indexTemplate
}
