package main

import (
	"errors"
	"go-restapi/mux-webservice/api"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	srv := api.NewServer()

	wd, _ := os.Getwd()
	cert_file := filepath.Join(wd, "/..", "localhost.crt")
	key_file := filepath.Join(wd, "/..", "localhost.key")

	use_https := true

	if _, err := os.Stat(cert_file); errors.Is(err, os.ErrNotExist) {
		use_https = false
	}
	if _, err := os.Stat(key_file); errors.Is(err, os.ErrNotExist) {
		use_https = false
	}

	if use_https {
		log.Fatal(http.ListenAndServeTLS(":9001", cert_file, key_file, srv))
	} else {
		http.ListenAndServe(":9000", srv)
	}

}
