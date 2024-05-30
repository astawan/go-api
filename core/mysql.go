package core

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLConnection(app *Application) *gorm.DB {
	// TODO import mysql url dari App
	db, err := gorm.Open(mysql.Open(app.Env.MysqlUrl))
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	return db
}
