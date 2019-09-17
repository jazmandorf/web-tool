package controlloer

import (
	_ "fmt"
	_ "net/http"

	"github.com/labstack/echo"
)

type LoginInfo struct {
	Email string
	Pass  string
}

func Login(c echo.Context) LoginInfo {
	loginInfo := LoginInfo{}
	loginInfo.Email = c.FormValue("email")
	loginInfo.Pass = c.FormValue("pass")
	return loginInfo
}
