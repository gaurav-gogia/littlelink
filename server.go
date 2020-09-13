package main

import (
	"fmt"
	"littlelink/dblayer"
	"littlelink/errrs"
	"littlelink/model"
	"littlelink/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/xujiajun/nutsdb"
)

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	db, err := dblayer.InitDB()
	errrs.Handle(err)
	defer db.Close()

	logFile, err := os.OpenFile(model.LOGSTRING, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	errrs.Handle(err)
	defer logFile.Close()

	log.SetOutput(logFile)

	baseRouter := router.NewBaseRouter(db)
	setupRouter(baseRouter)

	fmt.Println("Starting Server ....")
	go http.ListenAndServe(":"+model.LISTENPORT, nil)

	<-done
	clean(db, logFile)
}

func setupRouter(br *router.BaseRouter) {
	http.HandleFunc(model.CREATELITTLE, br.SetSmallRoute)
	http.HandleFunc(model.GETBIG, br.GetLongRoute)
}

func clean(db *nutsdb.DB, logFile *os.File) {
	errrs.Handle(db.Close())
	errrs.Handle(logFile.Close())
	fmt.Printf("\nShutting Down.\n")
}
