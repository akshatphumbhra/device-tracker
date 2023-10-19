package routes

import (
	"github.com/akshatphumbhra/device-tracker/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterDeviceRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/devices/", controllers.FetchDeviceData).Methods("GET")
	router.HandleFunc("/api/devices/update/visibility", controllers.UpdateDeviceVisibility).Methods("PATCH")
	router.HandleFunc("/api/devices/update/icon", controllers.UpdateDeviceIcon).Methods("PATCH")
}
