package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	echosession "github.com/go-session/echo-session"
	"github.com/labstack/echo"
)

type LoginInfo struct {
	Email string
	Pass  string
}

func MakeNameSpace(name string) string {
	now := time.Now()
	nanos := strconv.FormatInt(now.UnixNano(), 10)

	result := name + "-" + nanos
	fmt.Println("makeNameSpace : ", result)
	return result
}

func RegUserConrtoller(c echo.Context) error {
	user := c.FormValue("username")
	pass := c.FormValue("password")
	fmt.Println("c.Request : ", c.Request())
	store := echosession.FromContext(c)
	get, ok := store.Get(user)
	fmt.Println("GET store : ", get)
	obj := map[string]string{
		"username":  user,
		"namespace": MakeNameSpace(user),
		"password":  pass,
	}
	if !ok {

		store.Set(user, obj)
		err := store.Save()
		if err != nil {
			return c.JSON(http.StatusServiceUnavailable, map[string]string{
				"message": "Fail",
			})
		}
		return c.JSON(http.StatusOK, map[string]string{
			"message": "SUCCESS",
		})
	} else {
		return c.JSON(301, map[string]string{
			"message": "already register",
		})
	}

}

func LoginController(c echo.Context) error {

	getUser := c.FormValue("username")
	getPass := c.FormValue("password")
	store := echosession.FromContext(c)
	fmt.Println("c.Request : ", c.Request())

	fmt.Println(getUser, getPass)

	get, ok := store.Get(getUser)
	fmt.Println("getObj1 :", get)
	if !ok {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": " 정보가 없으니 다시 등록 해라",
		})
	}
	// var result map[string]string
	// for k, item := range getObj {
	// 	fmt.Println("GetItem : ", item)
	// 	result[k] = item
	// }
	fmt.Println("getObj :", get)
	// if sesEmail := session.Get(getUser); sesEmail != nil {
	// 	if sesEmail == getUser {

	// 	}
	// }

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success",
		"status":  "200",
		"user":    getUser,
		"pass":    getPass,
	})
}
