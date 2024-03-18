package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/c-128/transittracker/handlers"
	"github.com/c-128/transittracker/routes"
	"github.com/c-128/transittracker/services"
)

//go:embed views
var views embed.FS

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	templ, err := template.ParseFS(views, "views/*.html")
	if err != nil {
		return err
	}

	vehicles := services.NewVehicles()
	go updateVehicles(vehicles)

	mux := http.NewServeMux()

	mux.Handle("GET /", handlers.ErrorHandler(routes.GetOverview(templ, vehicles)))

	err = http.ListenAndServe("0.0.0.0:3000", mux)
	if err != nil {
		return err
	}

	return nil
}

func updateVehicles(vehicles *services.Vehicles) {
	ticker := time.NewTicker(5 * time.Minute)
	for {
		log.Printf("Fetching vehicles")

		err := vehicles.FetchVehicles()
		if err != nil {
			log.Printf("Failed to fetch vehicles: %s", err)
		}

		<-ticker.C
	}
}
