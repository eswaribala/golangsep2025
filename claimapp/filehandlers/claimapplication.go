package main

import (
	"fmt"
	"log"
	"time"

	"github.com/eswaribala/claimapp/interfaces"
	"github.com/eswaribala/claimapp/models"
	"github.com/eswaribala/claimapp/utility"
)

func main() {

	//interface
	var vehicleRepo interfaces.IVehicleRepo = nil

	//create 10 vehicle instances and save them to csv file
	for i := 1; i <= 10; i++ {
		vehicle := &models.Vehicle{
			RegistrationNo:     "KA-01-123" + fmt.Sprintf("%02d", i),
			Maker:              "Toyota",
			DateOfRegistration: time.Now(),
			ChassisNo:          "CH123456789",
			FuelType:           models.Petrol,
			EngineNo:           "EN123456789",
			Color:              "Red",
		}

		vehicleRepo = vehicle
		_, err := vehicleRepo.Save()
		if err != nil {
			log.Fatalf("Error saving vehicle: %v", err)
		}

	}

	log.Println("10 vehicles saved successfully.")
	vehicles, _ := vehicleRepo.GetAll()
	fmt.Println("All vehicles:")
	for _, v := range vehicles {
		fmt.Println("Vehicle:", v)
		for key, value := range utility.StructToMapVehicle(v) {

			fmt.Printf("%s : %v\n", key, value)
		}
	}

}
