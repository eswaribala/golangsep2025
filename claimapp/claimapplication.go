package main

import (
	"fmt"
	"time"

	"github.com/eswaribala/claimapp/interfaces"
	"github.com/eswaribala/claimapp/models"
	"github.com/eswaribala/claimapp/utility"
)

func main() {

	var vehicleRepo interfaces.IVehicleRepo = nil
	//vehicle instance
	vehicle := &models.Vehicle{
		RegistrationNo:     "KA-01-1234",
		Maker:              "Toyota",
		DateOfRegistration: time.Now(),
		ChassisNo:          "CH123456789",
		FuelType:           models.Petrol,
		EngineNo:           "EN123456789",
		Color:              "Red",
	}
	//assigning to interface
	vehicleRepo = vehicle
	result, _ := vehicleRepo.Save()
	println(result)
	data, error := vehicleRepo.GetByID(vehicle.RegistrationNo)
	if error != nil {
		println(error.Error())
	}
	for key, value := range utility.StructToMapVehicle(data) {

		fmt.Printf("%s : %v\n", key, value)
	}
	vehicles, _ := vehicleRepo.GetAll()
	fmt.Println("All vehicles:")
	for _, v := range vehicles {
		for key, value := range utility.StructToMapVehicle(v) {

			fmt.Printf("%s : %v\n", key, value)
		}
	}
	data, error = vehicleRepo.Update(vehicle.RegistrationNo, "Blue")
	if error != nil {
		println(error.Error())
	}
	for key, value := range utility.StructToMapVehicle(data) {

		fmt.Printf("%s : %v\n", key, value)
	}

	status, _ := vehicleRepo.Delete(vehicle.RegistrationNo)
	fmt.Println("Deleted:", status)

}
