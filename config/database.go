package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(env Env, logger Logger) (*Database, error) {
	username := env.DBUsername
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName
	dbsslmode := env.DBSslmode

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=America/Bogota", host, username, password, dbname, port, dbsslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.GetGormLogger(),
	})

	if err != nil {
		logger.Info("Url/dsn: ", dsn)
		logger.Panic(err)
	}

	logger.Info("Database connection established")

	return &Database{
		DB: db,
	}, nil
}

func (db *Database) CloseDBConnection() {
	sqlDB, _ := db.DB.DB()
	sqlDB.Close()
}
