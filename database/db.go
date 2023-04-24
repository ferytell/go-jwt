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
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"), os.Getenv("PGPORT"))
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}
	fmt.Println("sukses terkoneksi ke database")
	// Migrate models
	db.Debug().AutoMigrate(
		models.User{},
		models.Comments{},
		models.Photo{},
		models.SocialMedia{},
	)

	// 	// Create trigger
	// 	_, err := db.Exec(`CREATE OR REPLACE FUNCTION delete_comments()
	// 	RETURNS TRIGGER AS
	// 	$$
	// 	BEGIN
	// 		DELETE FROM comments WHERE photo_id = OLD.id;
	// 		RETURN OLD;
	// 	END;
	// 	$$ LANGUAGE plpgsql;
	// 	`),
	// 	if err != nil {
	// 	return err
	// 	}
	// 	//err = db.Exec(`CREATE TRIGGER delete_comments AFTER DELETE ON photos FOR EACH ROW EXECUTE FUNCTION delete_comments();`)
	// 	err = db.Exec(`CREATE TRIGGER delete_comments AFTER DELETE ON photos FOR EACH ROW BEGIN DELETE FROM comments WHERE photo_id = OLD.id; END;`).Error
	// 	if err != nil {
	// 		log.Fatal("error creating trigger: ", err)
	// 	}

	// 	fmt.Println("successfully created trigger")

}

func GetDB() *gorm.DB {
	return db
}
