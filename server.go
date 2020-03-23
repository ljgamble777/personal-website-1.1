package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", handlers.LoggingHandler(os.Stdout, fs))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
