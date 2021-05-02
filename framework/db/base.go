package db

import (
	"log"
	"os"

	"github.com/rodrigopmatias/ligistic/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type OpenConfig struct {
	Migrate bool
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		value = defaultValue
	}

	return value
}

func Open(config OpenConfig) (*gorm.DB, error) {
	var err error
	var conn *gorm.DB

	conn, err = database()

	if err != nil {
		log.Println("Can't connect with database")
		log.Fatalln(err)

		return nil, err
	}

	if config.Migrate {
		log.Println("Running migrations of models")
		conn.AutoMigrate(&models.User{})
	}

	return conn, nil
}

func database() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	switch dialect := getEnv("DB_DIALECT", "sqlite"); dialect {
	case "postgres":
		dbDSN := getEnv("DB_DSN", "host=127.0.0.1 port=5432 sslmode=disable user=postgres password=secret dbname=db")
		db, err = gorm.Open(postgres.Open(dbDSN), &gorm.Config{})
	case "mysql":
		dbDSN := getEnv("DB_DSN", "user:passwd@tcp(127.0.0.1:3306)/db?charset=utf8")
		db, err = gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
	case "sqlite":
		dbDSN := getEnv("DB_DSN", "db.sqlite3")
		db, err = gorm.Open(sqlite.Open(dbDSN), &gorm.Config{})
	}

	return db, err
}
