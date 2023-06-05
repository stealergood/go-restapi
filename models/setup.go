package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type Book struct {
	Id			int64	`gorm:"primaryKey" json:"id"`
	Title		string	`gorm:"type:varchar(255)" json:"title"`
	Description	string	`gorm:"type:text" json:"description"`
	Author		string	`gorm:"type:varchar(255)" json:"author"`
	PublishDate	string	`gorm:"type:date" json:"publish_date"`
}

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("denny:deni123@tcp(localhost:3306)/go_restapi_fiber"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Book{})
	DB = db
}