package main

import (
	"log"
	"net/http"

	"github.com/javatour/runserver/worker"
)

func main() {
	workers, err := worker.MakeWorkers()
	if err != nil {
		log.Fatal("do not use this program now. your server already busy")
	}
	workers.WorkStart()
	http.HandleFunc("/code", workers.ServeHTTP)
	log.Fatal(http.ListenAndServe(":3001", nil))
}
