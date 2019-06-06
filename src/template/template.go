package template

import (
	"html/template"
	"net/http"
)

const (
	// PathTemplate is the relative path to the directory storing
	// all the GoHTML files / templates
	PathTemplate = "./public/template"
)

// Init will create a new template manager
func Init() error {
	return nil
}

func loadSingleTemplate(templateName string) (*template.Template, error) {
	return template.New(templateName).ParseFiles(CalculPath(templateName))
}

// CalculPath returns the relative path of the given template name
func CalculPath(templateName string) string {
	return PathTemplate + "/" + templateName + ".gohtml"
}

// Validate the given template, return a non-nil error if the template is invalid
func Validate(templateName string) error {
	_, err := template.New("html-tmpl").ParseFiles(CalculPath(templateName))
	return err
}

// RenderTemplate renders the given templateName to the given
// http.responseWritter, data will be send to the template
func RenderTemplate(w http.ResponseWriter, data interface{}, templateName string) error {
	tmpl, err := loadSingleTemplate(templateName)
	if err != nil {
		return err
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}
