package repository

import (
	"rap/dbConfig"
	"rap/models"
	"rap/utils"
)

// Permissions : singleton
var Permissions []models.Permission

// GetPermissons find all permission by zone id
func GetPermissons(zoneID int) []models.Permission {
	if Permissions == nil {
		if err := dbConfig.DB.Where("zone_id = ?", zoneID).Find(&Permissions).Error; err != nil {
			utils.Check(err)
		}
	}
	return Permissions
}
