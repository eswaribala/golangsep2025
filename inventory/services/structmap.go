package services

import (
	"inventory/models"

	"github.com/mitchellh/mapstructure"
)

func StructToMapPatient(p *models.Patient) map[string]interface{} {
	var patientMap map[string]interface{}
	_ = mapstructure.Decode(p, &patientMap)
	return patientMap
}

func StructToMapDate(p *models.Date) map[string]interface{} {
	var dateMap map[string]interface{}
	_ = mapstructure.Decode(p, &dateMap)
	return dateMap
}
