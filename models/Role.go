package models

type Role struct {
	ID     int
	Name   string `gorm:"column:NAME"`
	ZoneID int    `gorm:"column:ZONE_ID"`
	Code   string `gorm:"column:CODE"`
	Status int    `gorm:"column:STATUS"`
}

func (Role) TableName() string {
	return "SE_ROLE"
}
