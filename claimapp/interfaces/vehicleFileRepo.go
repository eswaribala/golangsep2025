package interfaces

import "github.com/eswaribala/claimapp/models"

type IVehicleFileRepo interface {
	SaveToFile(fileName string, vehicles, headers []string, vehicleModels []*models.Vehicle) (bool, error)
}
