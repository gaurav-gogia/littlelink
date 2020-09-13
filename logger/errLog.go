package logger

import (
	"littlelink/model"
	"log"
	"net/http"
)

// ErrLog function logs an error to the log file
func ErrLog(severity int, err string, r *http.Request) {
	log.Println(model.LOGLINE)
	log.Printf("Route: %s\n", r.URL.Path)
	log.Printf("IP: %s\n", r.RemoteAddr)
	log.Printf("Severity: %d\n", severity)
	log.Printf("Error: %s\n", err)
	log.Printf("Method: %s\n", r.Method)
}
