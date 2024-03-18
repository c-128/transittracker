package bwegt_api

import (
	"encoding/json"
	"net/http"

	"github.com/c-128/transittracker/models"
)

func FetchVehicles() ([]models.Vehicle, error) {
	res, err := http.Get("https://www.efa-bw.de/VELOC?CoordSystem=WGS84")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var vehicles []models.Vehicle
	err = json.NewDecoder(res.Body).Decode(&vehicles)
	if err != nil {
		return nil, err
	}

	return vehicles, nil
}
