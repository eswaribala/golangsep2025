package models

import (
	"fmt"
	"time"
)

// custom type
type FuelType string

const (
	Petrol   FuelType = "petrol"
	Diesel   FuelType = "diesel"
	Electric FuelType = "electric"
)

type Vehicle struct {
	RegistrationNo     string    `json:"id"`
	Maker              string    `json:"make"`
	DateOfRegistration time.Time `json:"date_of_registration"`
	ChassisNo          string    `json:"chassis_no"`
	FuelType           FuelType  `json:"fuel_type"`
	EngineNo           string    `json:"engine_no"`
	Color              string    `json:"color"`
}

// create map
var vehicleMap = make(map[string]*Vehicle)

func (v *Vehicle) Save() (bool, error) {
	vehicleMap[v.RegistrationNo] = v
	return true, nil
}
func (v *Vehicle) GetByID(id string) (*Vehicle, error) {
	if vehicle, exists := vehicleMap[id]; exists {
		return vehicle, nil
	}
	return nil, fmt.Errorf("vehicle not found")
}
func (v *Vehicle) GetAll() ([]*Vehicle, error) {
	vehicles := make([]*Vehicle, 0, len(vehicleMap))
	for _, vehicle := range vehicleMap {
		vehicles = append(vehicles, vehicle)
	}
	return vehicles, nil
}
func (v *Vehicle) Update(id string, color string) (*Vehicle, error) {
	if vehicle, exists := vehicleMap[id]; exists {
		vehicle.Color = color
		return vehicle, nil
	}
	return nil, fmt.Errorf("vehicle not found")
}
func (v *Vehicle) Delete(id string) (bool, error) {
	if _, exists := vehicleMap[id]; exists {
		delete(vehicleMap, id)
		return true, nil
	}
	return false, fmt.Errorf("vehicle not found")
}
