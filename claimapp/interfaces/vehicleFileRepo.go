package interfaces

import "github.com/eswaribala/claimapp/models"

type IVehicleFileRepo interface {
	SaveToFile(fileName string, headers []string, vehicleModels []*models.Vehicle) (bool, error)
	SaveToMongoDB(vehicleModels []*models.Vehicle) (bool, error)
	SaveToPostgres(vehicleModels []*models.Vehicle) (bool, error)
}
