package routes

import (
	"fmt"
	sv "../../service"
)

var routes Route

routes = []Route{
	//----------CloudDriverInfo
	
		{"POST", "/driver", sv.RegisterCloudDriver},
		{"GET", "/driver", sv.ListCloudDriver},
		{"GET", "/driver/:DriverName", getCloudDriver},
		{"DELETE", "/driver/:DriverName", unRegisterCloudDriver},
	}