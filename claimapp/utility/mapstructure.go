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

func StructToMapLocation(p *models.Location) map[string]interface{} {
	var locationMap map[string]interface{}
	_ = mapstructure.Decode(p, &locationMap)
	return locationMap
}