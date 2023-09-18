package cookiemanager

import (
	"net/http"
	"time"
)

func CreateCookie(name string, value string, timeExp int) http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = time.Now().Add(time.Duration(timeExp) * time.Second)
	return *cookie
}
