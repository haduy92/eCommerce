package database

import (
	"eCommerce/infrastructure/errs"
	"eCommerce/model/entity"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}

func BuildDBConfig() *DBConfig {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	intPort, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		errs.E(errs.Internal, err)
	}

	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     intPort,
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
	return &dbConfig
}

func Init() *gorm.DB {

	dbConfig := BuildDBConfig()
	dbURL := DbURL(dbConfig)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err1 := db.AutoMigrate(&entity.Person{})
	if err1 != nil {
		panic(err)
	}
	return db
}

func Dispose(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.Close()
}
