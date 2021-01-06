package models

type API struct {
	ID     int	 `gorm:"primarykey"`
	API    string `gorm:"column:API"`
	Method string `gorm:"column:METHOD"`
	Status int    `gorm:"column:STATUS"`
	ZoneID int    `gorm:"column:ZONE_ID"`
}

func (API) TableName() string {
	return "API"
}
