package interfaces

import (
	"github.com/eswaribala/claimapp/models"
)

type IVehicleFileRepo interface {
	SaveToFile(vehicles []*models.Vehicle) (bool, error)
}
