package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a name for the migration.")
		return
	}

	migrationName := os.Args[1]
	timestamp := time.Now().Format("20060102150405")
	basePath := "internal/infrastructure/database/migrations"

	upFileName := fmt.Sprintf("%s/%s_%s.up.sql", basePath, timestamp, migrationName)
	downFileName := fmt.Sprintf("%s/%s_%s.down.sql", basePath, timestamp, migrationName)

	if err := createFile(upFileName); err != nil {
		fmt.Printf("Error creating migration up file: %v\n", err)
		return
	}

	if err := createFile(downFileName); err != nil {
		fmt.Printf("Error creating migration down file: %v\n", err)
		return
	}

	fmt.Printf("Migration files created:\n%s\n%s\n", upFileName, downFileName)
}

func createFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
