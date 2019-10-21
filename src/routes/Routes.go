package routes

import (
	"github.com/labstack/echo"
)

type Route struct {
	method, path string
	function     echo.HandlerFunc
}
