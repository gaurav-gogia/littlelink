package logger

import (
	"littlelink/model"
	"log"
	"net/http"
	"time"
)

// CallLog function logs the route, method and time it took to complete that API call
func CallLog(r *http.Request, time time.Duration) {
	log.Println(model.LOGLINE)
	log.Printf("Route: %s\n", r.URL.Path)
	log.Printf("IP: %s\n", r.RemoteAddr)
	log.Printf("Method: %s\n", r.Method)
	log.Printf("Function Time: %v\n", time)
}
