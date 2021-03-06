package template

import (
	"github.com/pedafy/pedafy/src/server/user"
	"google.golang.org/appengine"
)

// Page represents a data structure used to render information
// on a Go template
type Page struct {
	PageName   string
	Production bool
	LoggedIn   bool
	User       user.User
	Data       interface{}
}

// NewPage returns a new page composed of the given data
func NewPage(pageName string, loggedIn bool, user user.User, data interface{}) Page {
	return Page{PageName: pageName,
		Data:       data,
		Production: !appengine.IsDevAppServer(),
		LoggedIn:   loggedIn,
		User:       user,
	}
}
