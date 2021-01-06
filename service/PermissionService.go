package service

import (
	"fmt"
	"rap/repository"
)

// GenPermissionInsertStatement 2 params
// zoneID : current zone id
// newZoneID : new zone id
func GenPermissionInsertStatement(zoneID int, newZoneID int) {
	statement += "\n"
	permissions := repository.GetPermissons(zoneID)
	for _, permission := range permissions {
		statement += fmt.Sprintf("insert into auth.PERMISSION (`NAME` , zone_id, `status`) values ('%s', %d, %d);\n", permission.Name, newZoneID, permission.Status)
	}
}
