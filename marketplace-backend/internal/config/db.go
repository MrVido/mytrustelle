package config

import (
    "fmt"
    "log"
    "os"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

// DBConfig represents the database configuration.
type DBConfig struct {
    Host     string
    Port     int
    User     string
    DBName   string
    Password string
    SSLMode  string
}

// BuildDBConfig constructs the DBConfig.
func BuildDBConfig() *DBConfig {
    dbConfig := DBConfig{
        Host:     "localhost",
        Port:     5432,
        User:     "your_database_user",
        Password: "your_database_password",
        DBName:   "your_database_name",
        SSLMode:  "disable", // Change to "require" if your environment uses SSL.
    }
    return &dbConfig
}

// DbURL constructs the database connection string from DBConfig.
func DbURL(dbConfig *DBConfig) string {
    return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
        dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.DBName, dbConfig.SSLMode, dbConfig.Password)
}

// ConnectDatabase initializes the database connection using GORM.
func ConnectDatabase() *gorm.DB {
    dbConfig := BuildDBConfig()
    db, err := gorm.Open(postgres.Open(DbURL(dbConfig)), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    fmt.Println("Successfully connected to the database.")
    return db
}
