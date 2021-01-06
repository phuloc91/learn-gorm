package service

import (
	"fmt"
	"rap/repository"
)

// GenRoleInsertStatement 2 params
// zoneID : current zone id
// newZoneID : new zone id
func GenRoleInsertStatement(zoneID int, newZoneID int) {
	statement += "\n"
	roles := repository.GetRoles(zoneID)
	for _, role := range roles {
		statement += fmt.Sprintf("insert into auth.SE_ROLE (`CODE`, `NAME`, zone_id, `status`) values ('%s', '%s', %d, %d);\n", role.Code, role.Name, newZoneID, role.Status)
	}
}
