package main

import (
	"github.com/TeamTells/portal_documents_service/api"
	"github.com/TeamTells/portal_documents_service/config"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {

	// read config from env
	cfg := config.Read()

	err := http.ListenAndServe(cfg.HTTPAddr+":"+cfg.HTTPPort, api.Handler(&api.Implementation{}))
	if err != nil {
		return err
	}

	return nil
}
