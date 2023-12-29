package configs

import (
	"jti-test/internal/models"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func InitializeDatabase() {
	dsn := "host=localhost user=postgres password=password dbname=phonenumber port=5432 sslmode=disable" // Adjust connection string as needed
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	tx := db.Session(&gorm.Session{PrepareStmt: true})

	if err != nil {
		log.Fatal("Koneksi DB Gagal")
	}

	migrateDatabase(tx)
}

func GetDB() *gorm.DB {
	return db
}

func migrateDatabase(db *gorm.DB) {
	if !fiber.IsChild() {

		errMigrate := db.AutoMigrate(&models.PhoneNumber{}, &models.Provider{})

		if errMigrate != nil {
			log.Fatal("Gagal Migrate")
		}

		log.Println("Migrate Berhasil!")
	}

}
