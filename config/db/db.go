package config

import (
	"fmt"
	env "go_project_structure/config/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() (*gorm.DB, error) {
	host := env.GetString("DB_HOST", "127.0.0.1")
	port := env.GetString("DB_PORT", "5432")
	user := env.GetString("DB_USER", "minhaz_hossain")
	password := env.GetString("DB_PASSWORD", "12345")
	dbname := env.GetString("DB_NAME", "auth_dev")
	sslmode := env.GetString("DB_SSLMODE", "disable")
	timezone := env.GetString("DB_TIMEZONE", "UTC")

	// Example of using these values to construct a connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone,
	)

	// fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return nil, err
	}

	pgsqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Failed to get database connection:", err)
		return nil, err
	}
	err = pgsqlDB.Ping()
	if err != nil {
		fmt.Println("Failed to ping database:", err)
		return nil, err
	}

	fmt.Println("Successfully connected to database")

	var dbName string
	db.Raw("SELECT current_database()").Scan(&dbName)
	fmt.Println("Connected to database:", dbName)

	return db, nil

}


/*
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


INSERT INTO users (name, email, password) VALUES ('John Doe', 'john@example.com', 'password123');

SELECT id, name, email FROM users;
*/
