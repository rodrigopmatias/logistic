package db

import (
	"log"
	"os"

	"github.com/rodrigopmatias/ligistic/framework/config"
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
		conn.AutoMigrate(
			&models.User{},
			&models.Wallet{},
			&models.Category{},
		)
	}

	return conn, nil
}

func database() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	cnf := config.New()

	switch cnf.DbDialect {
	case "postgres":
		db, err = gorm.Open(postgres.Open(cnf.DbDSN), &gorm.Config{})
	case "mysql":
		db, err = gorm.Open(mysql.Open(cnf.DbDSN), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(cnf.DbDSN), &gorm.Config{})
	}

	return db, err
}
