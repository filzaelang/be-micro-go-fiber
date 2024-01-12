package models

import "gorm.io/gorm"

type Article struct {
	ID      uint    `gorm:"column:id;primary_key;autoIncrement" json:"id"`
	Title   *string `gorm:"column:title" json:"title"`
	Author  *string `gorm:"column:author" json:"author"`
	Article *string `gorm:"column:article" json:"article"`
	Image   *string `gorm:"column:image" json:"image"`
	Date    *string `gorm:"column:date" json:"date"`
	UserId  *int    `gorm:"column:userId" json:"userId"`
}

func MigrateArticles(db *gorm.DB) error {
	err := db.AutoMigrate(&Article{})
	return err
}
