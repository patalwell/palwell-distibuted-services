package main

import (
	"log"
	"patalwell/palwell-distributed-services/internal/server"
)

func main() {
	srv := server.NewHttpServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
