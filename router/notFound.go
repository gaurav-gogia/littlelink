package router

import "net/http"

// NotFound funciton is a drop in API for all the routes NOT defined in the router
// Its job is to serve a good response for API consumers
func NotFound(w http.ResponseWriter, r *http.Request) {
	// TODO: Complete this function
}
