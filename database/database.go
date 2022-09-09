package database

import (
	"GORM-Practice/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func NewConnection() (*gorm.DB, error) {

	err := godotenv.Load("database/.env")

	if err != nil {
		log.Fatal(err)
	}

	config := &Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return DB, err

}

func Migrate() error {
	// db.Migrator().DropTable(&models.User{}, &models.Company{})
	// db.Migrator().DropTable(&models.User{}, &models.Company{}, &models.Job{}, &models.Applicant{})
	// err := db.AutoMigrate(&models.User{}, &models.Company{}, &models.Job{}, &models.Applicant{})
	DB.Migrator().DropTable(&models.Faculty{}, &models.Student{})

	err := DB.AutoMigrate(&models.Faculty{}, &models.Student{})
	return err
}

func LoadDB() {
	DB, _ = NewConnection()
	Migrate()

}
