package routes

import (
	"html/template"
	"net/http"

	"github.com/c-128/transittracker/handlers"
	"github.com/c-128/transittracker/services"
)

func GetOverview(templ *template.Template, vehicles *services.Vehicles) handlers.Error {
	return func(writer http.ResponseWriter, req *http.Request) error {
		vehicles, err := vehicles.GetVehicles()
		if err != nil {
			return err
		}

		/*
			0 - Zug
			1 - Zug
			2 -
			3 - Zug
			4 - Straßenbahn
			5 - Bus
			6 - Bus
			7 -
			8 - Zug
			9 - Fähre
			10 - Rufbus
			11 -
			12 -
			13 - Zug
			14 - Zug
			15 - Zug
			16 - Zug
		*/

		totalVehicles := 0
		totalVehiclesTrains := 0
		totalVehiclesTrams := 0
		totalVehiclesBuses := 0
		totalVehiclesFerries := 0
		totalVehiclesOnDemandBuses := 0
		totalVehiclesOther := 0

		realtimeAvailable := 0
		realtimeUnavailable := 0
		realtimeUnkown := 0

		delayTotal := 0
		delayDelayed := 0
		delayNotDelayed := 0

		for _, vehicle := range vehicles {
			totalVehicles++

			switch vehicle.MOTCode {
			case 0:
				totalVehiclesTrains++
			case 1:
				totalVehiclesTrains++
			case 3:
				totalVehiclesTrains++
			case 4:
				totalVehiclesTrams++
			case 5:
				totalVehiclesBuses++
			case 6:
				totalVehiclesBuses++
			case 8:
				totalVehiclesTrains++
			case 9:
				totalVehiclesFerries++
			case 10:
				totalVehiclesOnDemandBuses++
			case 13:
				totalVehiclesTrains++
			case 14:
				totalVehiclesTrains++
			case 15:
				totalVehiclesTrains++
			case 16:
				totalVehiclesTrains++
			default:
				totalVehiclesOther++
			}

			switch vehicle.RealtimeAvailable {
			case 1:
				realtimeAvailable++
			case 0:
				realtimeUnavailable++
			default:
				realtimeUnkown++
			}

			delayTotal += vehicle.Delay / 60
			if vehicle.Delay > 0 {
				delayDelayed++
			} else {
				delayNotDelayed++
			}
		}

		err = templ.ExecuteTemplate(writer, "overview.html", handlers.Map{
			"Vehicles": handlers.Map{
				"Total":         totalVehicles,
				"Trains":        totalVehiclesTrains,
				"Trams":         totalVehiclesTrams,
				"Buses":         totalVehiclesBuses,
				"Ferries":       totalVehiclesFerries,
				"OnDemandBuses": totalVehiclesOnDemandBuses,
				"Other":         totalVehiclesOther,
			},
			"Realtime": handlers.Map{
				"Available":   realtimeAvailable,
				"Unavailable": realtimeUnavailable,
				"Unkown":      realtimeUnkown,
			},
			"Delay": handlers.Map{
				"Total":      delayTotal,
				"Delayed":    delayDelayed,
				"NotDelayed": delayNotDelayed,
			},
		})
		if err != nil {
			return err
		}

		return nil
	}
}
