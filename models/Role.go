package models

type Role struct {
	ID          int 	`gorm:"primarykey"`
	Name        string        `gorm:"column:NAME"`
	ZoneID      int           `gorm:"column:ZONE_ID"`
	Code        string        `gorm:"column:CODE"`
	Status      int           `gorm:"column:STATUS"`
	Permissions []*Permission `gorm:"many2many:ROLE_PERMISSION;"`
}

func (Role) TableName() string {
	return "SE_ROLE"
}
