package service

import (
	"log"
	"net/http"
)

func StartWebServer(port string) {
	r := NewRouter()
	http.Handle("/", r)

	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil) //Goroutine will block here

	if err != nil {
		log.Println("An error occured starting HTTP listener")
		log.Println("Error: " + err.Error())
	}
}
