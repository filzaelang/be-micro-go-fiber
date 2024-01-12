package models

import "gorm.io/gorm"

type Partai struct {
	ID               uint    `gorm:"column:id;primary_key;autoIncrement" json:"id"`
	Name             *string `gorm:"column:name" json:"name"`
	Chairman         *string `gorm:"column:chairman" json:"chairman"`
	VisionAndMission *string `gorm:"column:visionAndMission" json:"visionAndMission"`
	Address          *string `gorm:"column:address" json:"address"`
	Image            *string `gorm:"column:image" json:"image"`
	SelectedPaslonId *int    `gorm:"column:selectedPaslonId" json:"selectedPaslonId"`
}

func MigratePartai(db *gorm.DB) error {
	err := db.AutoMigrate(&Partai{})
	return err
}
