package router

import (
	"fmt"
	"littlelink/logger"
	"net/http"
	"time"
)

// Index page serves as entry point and instructions guide for APIs
func (br *BaseRouter) Index(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	response := fmt.Sprintf(
		`
		Welcome to project LittleLink.
		Total API Routes
		1. http://localhost:8080/setSmall
		2. http://localhost:8080/getLong

		Working API Route:
		http://localhost:8080/setSmall

		How to use:
		http://localhost:8080/setSmall
		
		Make a POST call with following params in body:
		link : <working url>
	`)
	w.Write([]byte(response))

	go logger.CallLog(r, time.Since(start))
}
