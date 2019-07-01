package user

import (
	"errors"
	"net/http"
	"time"

	"github.com/markbates/goth"
)

type Login = string

var users map[Login]User
var sessions map[string]Login

type OAuth struct {
	Token        string
	TokenSecret  string
	TokenRefresh string
	Expire       *time.Time
}

type User struct {
	Login     string
	Firstname string
	Lastname  string
	Oauth     OAuth
}

// Init the stores
func Init() {
	users = map[Login]User{}
	sessions = map[string]Login{}
}

// NewUser stores the given goth.User into the store
func NewUser(w http.ResponseWriter, r *http.Request, user goth.User) error {
	var newUser User

	newUser.Login = user.Email
	newUser.Firstname = user.FirstName
	newUser.Lastname = user.LastName
	newUser.Oauth.Expire = &user.ExpiresAt
	newUser.Oauth.Token = user.AccessToken
	newUser.Oauth.TokenSecret = user.AccessTokenSecret
	newUser.Oauth.TokenRefresh = user.RefreshToken

	sessions[user.AccessToken] = user.Email
	users[user.Email] = newUser

	c := &http.Cookie{
		Name:    "session",
		Path:    "/",
		Value:   user.AccessToken,
		Expires: time.Now().Add(time.Hour * 24),
	}
	http.SetCookie(w, c)

	return nil
}

func LogoutUser(w http.ResponseWriter, r *http.Request) error {
	c, err := r.Cookie("session")
	if err != nil {
		return err
	}

	if _, ok := sessions[c.Value]; !ok {
		return errors.New("no user found")
	}

	users[sessions[c.Value]] = User{}
	sessions[c.Value] = ""

	c = &http.Cookie{
		Name:   "session",
		Path:   "/",
		Value:  "",
		MaxAge: 0,
	}
	http.SetCookie(w, c)

	return nil
}

// GetUser returns the user link to the request
func GetUser(r *http.Request) (user User, err error) {
	c, err := r.Cookie("session")
	if err != nil {
		return User{}, err
	}

	if login, ok := sessions[c.Value]; ok {
		return users[login], nil
	}
	return User{}, errors.New("no user found")
}
