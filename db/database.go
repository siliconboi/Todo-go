package db

import (
	"Appointy/Tasks/todo/structs"
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	sqlDB, _ := sql.Open("pgx", os.Getenv("DB_NAME"))

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	err = gormDB.AutoMigrate(&structs.Tasks{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("okay")
	return gormDB, err
}

type DBhandler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) DBhandler {
	return DBhandler{db}
}
