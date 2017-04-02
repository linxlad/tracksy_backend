package main

import (
	"net/http"
	r "gopkg.in/gorethink/gorethink.v3"
	"log"
	"github.com/linxlad/TraskyEA/events"
)

func main() {
	session, err := r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
		Database: "tracksy",
	})

	if err != nil {
		log.Panic(err.Error())
	}

	router := NewRouter(session)
	router.Handle(events.EARLY_ACCESS, addInterest)

	http.Handle("/", router)
	http.ListenAndServe(":4001", nil)
}