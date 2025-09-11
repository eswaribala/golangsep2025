package utility
import (
	"github.com/mitchellh/mapstructure"
	"github.com/eswaribala/claimapp/models"
)

func StructToMapVehicle(p *models.Vehicle) map[string]interface{} {
	var vehicleMap map[string]interface{}
	_ = mapstructure.Decode(p, &vehicleMap)
	return vehicleMap
}
