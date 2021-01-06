package models

type Permission struct {
	ID     int 	`gorm:"primarykey"`
	Name   string  `gorm:"column:NAME"`
	ZoneID int     `gorm:"column:ZONE_ID"`
	Status int     `gorm:"column:STATUS"`
	Roles  []Role `gorm:"many2many:ROLE_PERMISSION;"`
	APIS   []API  `gorm:"many2many:API_PERMISSION;"`
}

func (Permission) TableName() string {
	return "PERMISSION"
}
