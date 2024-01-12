package models

import "gorm.io/gorm"

type Voter struct {
	ID            uint    `gorm:"column:id;primary_key;autoIncrement" json:"id"`
	UserId        *string `gorm:"column:userId" json:"userId"`
	VotedPaslonId *string `gorm:"column:votedPaslonId" json:"votedPaslonId"`
}

func MigrateVoter(db *gorm.DB) error {
	err := db.AutoMigrate(&Voter{})
	return err
}
