package main

import (
	"net/http"

	"github.com/getgiddy/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	log.Info("Starting server on port 8000")

	log.Info(`
   __________     ___    ____  ____
  / ____/ __ \   /   |  / __ \/  _/
 / / __/ / / /  / /| | / /_/ // /  
/ /_/ / /_/ /  / ___ |/ ____// /   
\____/\____/  /_/  |_/_/   /___/   `)

	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Error(err)
	}
}
