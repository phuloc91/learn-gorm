package repository

import (
	"rap/dbConfig"
	"rap/models"
	"rap/utils"
)

// APIS singleton
var APIS []models.API

// GetAPIS by zone id
func GetAPIS(zoneID int) []models.API {
	if APIS == nil {
		if err := dbConfig.DB.Where("zone_id = ?", zoneID).Find(&APIS).Error; err != nil {
			utils.Check(err)
		}
	}
	return APIS
}
