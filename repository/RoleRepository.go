package repository

import (
	"rap/dbConfig"
	"rap/models"
	"rap/utils"
)

// Roles singleton
var Roles []models.Role

// GetRoles find all roles by zone id
func GetRoles(zoneID int) []models.Role {
	if Roles == nil {
		if err := dbConfig.DB.Where("zone_id = ?", zoneID).Find(&Roles).Error; err != nil {
			utils.Check(err)
		}
	}
	return Roles
}
