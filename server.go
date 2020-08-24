package main

import (
	"fmt"
	"littlelink/dblayer"
	"littlelink/errrs"
	"littlelink/model"
	"littlelink/router"
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

	setupRouter()
	fmt.Println("Starting Server ....")
	go http.ListenAndServe(":"+model.LISTENPORT, nil)

	<-done
	clean(db)
}

func setupRouter() {
	http.HandleFunc(model.CREATELITTLE, router.CreateLittle)
	http.HandleFunc(model.GETBIG, router.GetBig)
}

func clean(db *nutsdb.DB) {
	errrs.Handle(db.Close())
	fmt.Printf("\nShutting Down.\n")
}
