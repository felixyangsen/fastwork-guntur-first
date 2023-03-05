package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	db  *gorm.DB
	err error
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
)

func ConnectDB() {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: InitLog(),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(150)
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database")
}

func GetDB() *gorm.DB {
	return db
}
