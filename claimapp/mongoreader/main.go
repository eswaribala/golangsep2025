package main

import (
	"fmt"
	"log"

	"github.com/eswaribala/claimapp/models"
)

func main() {

	//var vehicleRepo interfaces.IVehicleRepo = nil

	// Fetch vehicles from MongoDB
	vehicles, err := models.FetchVehiclesFromMongoDB()
	if err != nil {
		log.Fatalf("Error fetching vehicles from MongoDB: %v", err)
	}

	// Process the fetched vehicles
	for _, v := range vehicles {
		fmt.Println("Fetched Vehicle:", v)
	}
}
