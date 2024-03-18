package services

import (
	"github.com/c-128/transittracker/bwegt_api"
	"github.com/c-128/transittracker/models"
)

func NewVehicles() *Vehicles {
	return &Vehicles{
		fetching: false,
		vehicles: nil,
	}
}

type Vehicles struct {
	fetching bool
	vehicles []models.Vehicle
}

func (v *Vehicles) GetVehicles() ([]models.Vehicle, error) {
	if v.vehicles == nil {
		err := v.FetchVehicles()
		if err != nil {
			return nil, err
		}
	}

	return v.vehicles, nil
}

func (v *Vehicles) FetchVehicles() error {
	if v.fetching {
		return nil
	}

	v.fetching = true
	defer func() {
		v.fetching = false
	}()

	vehicles, err := bwegt_api.FetchVehicles()
	if err != nil {
		return err
	}

	v.vehicles = vehicles
	return nil
}
