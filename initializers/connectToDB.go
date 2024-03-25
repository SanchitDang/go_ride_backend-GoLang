package initializers

import (
	"log"
	// "os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	// dsn := os.Getenv("DB_URL")
	dsn := "host=flora.db.elephantsql.com user=kbolycjd password=1V_Ul0lZPdJ_fSFGJCvj2MU94p4RVqww dbname=kbolycjd port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatal("Failed to connect to database")
	}

}
