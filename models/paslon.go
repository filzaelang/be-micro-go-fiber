package models

import "gorm.io/gorm"

type Paslon struct {
	ID               uint    `gorm:"column:id;primary_key;autoIncrement" json:"id"`
	Name             *string `gorm:"column:name" json:"name"`
	No               *string `gorm:"column:no" json:"no"`
	VisionAndMission *string `gorm:"column:visionAndMission" json:"visionAndMission"`
	Image            *string `gorm:"column:image" json:"image"`
}

func MigratePaslon(db *gorm.DB) error {
	err := db.AutoMigrate(&Paslon{})
	return err
}
