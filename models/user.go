package models

import "gorm.io/gorm"

type User struct {
	ID       uint    `gorm:"primary key;autoIncrement" json:"id"`
	Fullname *string `gorm:"column:fullname" json:"fullname"`
	Address  *string `gorm:"column:address" json:"address"`
	Gender   *string `gorm:"column:gender" json:"gender"`
	Username *string `gorm:"column:username" json:"username"`
	Password *string `gorm:"column:password" json:"password"`
	Age      *int    `gorm:"column:age" json:"age"`
	Role     *string `gorm:"column:role" json:"role"`
}

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}
