package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"golang.org/x/crypto/acme/autocert"
)

var staticDir = flag.String("static", "./static", "directory where static files are stored")
var config = flag.String("config", "dev", "configuration to run (dev or prod)")

func main() {
	flag.Parse()

	fs := http.FileServer(http.Dir(*staticDir))
	http.Handle("/", handlers.LoggingHandler(os.Stdout, fs))

	if *config == "dev" {
		log.Fatal(http.ListenAndServe(":80", nil))
	} else if *config == "prod" {
		certManager := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist("henrywfisher.com", "www.henrywfisher.com"),
			Cache:      autocert.DirCache("."),
		}

		server := &http.Server{
			Addr: ":https",
			TLSConfig: &tls.Config{
				GetCertificate: certManager.GetCertificate,
			},
		}

		go http.ListenAndServe(":http", handlers.LoggingHandler(os.Stdout, certManager.HTTPHandler(nil)))

		log.Fatal(server.ListenAndServeTLS("", ""))
	} else {
		log.Fatal("invalid configuration")
	}
}
