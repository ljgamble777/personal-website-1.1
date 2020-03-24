# Personal Website

This code runs my website: [https://henrywfisher.com](https://henrywfisher.com)

Components
- The web server is written in Go using the [net/http](https://golang.org/pkg/net/http/) package.
- The web server runs on an AWS EC2 instance as a daemon using systemd.
- https is supported through the [autocert](https://pkg.go.dev/golang.org/x/crypto/acme/autocert?tab=doc) package, which uses Let's Encrypt.
- The front end is HTML/CSS with a tiny bit of javascript. Probably not compatible with some early versions of Internet Explorer.
- The deploy script builds the web server, stops the daemon, copies the binary and static files using rsync, and then starts the server.

The server uses http by default so that local testing is easier.