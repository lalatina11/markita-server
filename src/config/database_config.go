package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
	Timezone string
}

func NewDatabaseConfig() *DatabaseConfig {
	Host := GetEnv("DATABASE_HOST")
	User := GetEnv("DATABASE_USER")
	Password := GetEnv("DATABASE_PASSWORD")
	Name := GetEnv("DATABASE_NAME")
	Port := GetEnv("DATABASE_PORT")
	Timezone := GetEnv("DATABASE_TIME_ZONE")

	return &DatabaseConfig{Host, User, Password, Name, Port, Timezone}
}

func (this *DatabaseConfig) ToDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", this.Host, this.User, this.Password, this.Name, this.Port, this.Timezone)
}

func (this *DatabaseConfig) Connect() *gorm.DB {
	dsn := this.ToDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DATABASE CONNECTION FAILED")
	}
	return db
}
