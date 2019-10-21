package service

import (

	// REST API (echo)
	"net/http"

	"github.com/labstack/echo"
)

//================ CloudDriver Handler
func RegisterCloudDriver(c echo.Context) error {
	cblog.Info("call registerCloudDriver()")

	req := &dim.CloudDriverInfo{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	cldinfoList, err := dim.RegisterCloudDriverInfo(*req)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, &cldinfoList)
}

func ListCloudDriver(c echo.Context) error {
	cblog.Info("call listCloudDriver()")

	cldinfoList, err := dim.ListCloudDriver()
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, &cldinfoList)
}

func GetCloudDriver(c echo.Context) error {
	cblog.Info("call getCloudDriver()")

	cldinfo, err := dim.GetCloudDriver(c.Param("DriverName"))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, &cldinfo)
}

func UnRegisterCloudDriver(c echo.Context) error {
	cblog.Info("call unRegisterCloudDriver()")

	result, err := dim.UnRegisterCloudDriver(c.Param("DriverName"))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, &result)
}
