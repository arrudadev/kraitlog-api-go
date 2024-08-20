package main

import (
	"log"
	"os"

	"github.com/arrudadev/kraitlog-api/config"
	"github.com/arrudadev/kraitlog-api/internal/infrastructure/database"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := database.NewConnection(config.DBMigrationHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Error creating migration driver: %v", err)
	}

	migrationsPath := "file://internal/infrastructure/database/migrations"
	migration, err := migrate.NewWithDatabaseInstance(migrationsPath, "postgres", driver)
	if err != nil {
		log.Fatalf("Error creating migration instance: %v", err)
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "up":
			if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
				log.Fatalf("Error applying migrations: %v", err)
			}
			log.Println("Migrations applied successfully!")
		case "down":
			if err := migration.Down(); err != nil && err != migrate.ErrNoChange {
				log.Fatalf("Error reverting migrations: %v", err)
			}
			log.Println("Migrations reverted successfully!")
		default:
			log.Fatalf("Unknown command: %s", os.Args[1])
		}
	} else {
		log.Println("Please specify 'up' or 'down' as an argument.")
	}
}
