package database

import (
	"fmt"
	"log"
	"os"

	"github.com/ferytell/go-jwt/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	//host     = "xxx"
	//user     = "xxx"
	//password = "xxx"
	//dbPort   = "xxx"
	//dbName   = "xxx"
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("USER"), os.Getenv("PASS"), os.Getenv("DBNAME"), os.Getenv("DBPORT"))
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}
	fmt.Println(os.Getenv("HOST"))
	fmt.Println("sukses terkoneksi ke database")
	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
