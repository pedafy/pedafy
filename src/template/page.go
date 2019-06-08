package template

// Page represents a data structure used to render information
// on a Go template
type Page struct {
	PageName string
	Data     interface{}
}

// NewPage returns a new page composed of the given data
func NewPage(pageName string, data interface{}) Page {
	return Page{PageName: pageName,
		Data: data,
	}
}
