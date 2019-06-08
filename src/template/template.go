package template

import (
	"html/template"
	"net/http"
)

var (
	// pathTemplate is the relative path to the directory storing
	// all the GoHTML files / templates
	pathTemplate string

	tmpl *template.Template
)

// Init will create a new template manager
func Init(templatePath string) error {
	var err error
	pathTemplate = templatePath
	tmpl, err = loadTemplates()
	return err
}

func loadTemplates() (*template.Template, error) {
	return template.ParseGlob(pathTemplate + "/*.gohtml")
}

// CalculPath returns the relative path of the given template name
func CalculPath(templateName string) string {
	return pathTemplate + "/" + templateName + ".gohtml"
}

// RenderTemplate renders the given templateName to the given
// http.responseWritter, data will be send to the template
func RenderTemplate(w http.ResponseWriter, page Page, templateName string) error {
	err := tmpl.ExecuteTemplate(w, templateName, page)
	if err != nil {
		return err
	}
	return nil
}
