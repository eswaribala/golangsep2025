package interfaces

import "github.com/eswaribala/claimapp/models"

type IVehicleRepo interface {
	Save(vehicle models.Vehicle) (bool, error)
	GetByID(id string) (models.Vehicle, error)
	GetAll() ([]models.Vehicle, error)
	Update(id string, color string) (models.Vehicle, error)
	Delete(id string) (bool, error)
}
