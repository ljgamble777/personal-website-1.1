package main

import (
	"crypto/tls"
	"embed"
	"flag"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"golang.org/x/crypto/acme/autocert"
)

//go:embed static/*
var static embed.FS

var staticDir = flag.String("static", "./static", "directory where static files are stored")
var config = flag.String("config", "dev", "configuration to run (dev or prod)")

const localPort = ":80"

func main() {
	flag.Parse()

	root, err := fs.Sub(static, "static")
	if err != nil {
		log.Fatalf("could not get sub static directory: %v\n", err)
	}

	fs := http.FileServer(http.FS(root))
	http.Handle("/", handlers.LoggingHandler(os.Stdout, fs))

	if *config == "dev" {
		log.Printf("running at http://localhost%v", localPort)
		log.Fatal(http.ListenAndServe(localPort, nil))
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
