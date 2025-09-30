package session

import (
	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore
var SessionName string