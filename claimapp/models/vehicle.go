package models
//custom type
type FuelType string

const (
	Petrol FuelType = "petrol"
	Diesel FuelType = "diesel"
	Electric FuelType = "electric"
)

type Vehicle struct {
	RegistrationNo string   `json:"id"`
	Maker         string `json:"make"`
	ChassisNo        string `json:"chassis_no"`
	FuelType        FuelType `json:"fuel_type"`
	EngineNo        string `json:"engine_no"`
	Color        string `json:"color"`
}
