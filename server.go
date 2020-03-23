package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

var crtPath = flag.String("crt", "server.crt", "certificate location")
var keyPath = flag.String("key", "server.key", "key location")
var staticDir = flag.String("static", "./static", "directory where static files are stored")

func main() {
	flag.Parse()

	fs := http.FileServer(http.Dir(*staticDir))
	http.Handle("/", handlers.LoggingHandler(os.Stdout, fs))
	log.Fatal(http.ListenAndServe(":4000", nil))
	// log.Fatal(http.ListenAndServeTLS(":4000", *crtPath, *keyPath, nil))
}
