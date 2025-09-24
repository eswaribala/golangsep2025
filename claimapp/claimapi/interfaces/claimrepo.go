package interfaces

import "github.com/eswaribala/claimapp/claimapi/models"

type IClaimRepo interface {
	Save() (bool, error)
	GetByID(id uint) (*models.Claim, error)
	GetAll() ([]*models.Claim, error)
	Update(claimID uint, claimAmount int) (*models.Claim, error)
	Delete(id uint) (bool, error)
}
