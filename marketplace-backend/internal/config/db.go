package config

import (
    "fmt"
    "log"
    "os"
    "strconv"
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

// BuildDBConfig constructs the DBConfig using environment variables.
func BuildDBConfig() *DBConfig {
    port, err := strconv.Atoi(os.Getenv("DB_PORT"))
    if err != nil {
        log.Fatalf("Invalid DB Port: %v", err)
    }

    dbConfig := DBConfig{
        Host:     os.Getenv("DB_HOST"),
        Port:     port,
        User:     os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        DBName:   os.Getenv("DB_NAME"),
        SSLMode:  os.Getenv("DB_SSLMODE"),
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
    
    // Example of setting connection pool settings (adjust according to your application's needs)
    sqlDB, err := db.DB()
    if err != nil {
        log.Fatalf("Failed to get database: %v", err)
    }
    sqlDB.SetMaxIdleConns(10) // Set max idle connections
    sqlDB.SetMaxOpenConns(100) // Set max open connections
    sqlDB.SetConnMaxLifetime(1 * time.Hour) // Set the max lifetime of a connection

    fmt.Println("Successfully connected to the database.")
    return db
}
