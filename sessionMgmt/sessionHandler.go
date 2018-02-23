package sessionMgmt

import (
	"github.com/labstack/echo"
)

func LoginHandler(c echo.Context) {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if vefiry(username, password) {

	}
}

func vefiry(username, password string) bool {
	return true
}
