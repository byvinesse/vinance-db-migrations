package generate

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Generate() {
	// Parse command line arguments
	flag.Parse()

	// Get migration name from args
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Error: Migration name is required")
		fmt.Println("Usage: go run main.go generate <migration_name>")
		fmt.Println("Example: go run main.go generate create_users_table")
		os.Exit(1)
	}

	fmt.Println(args)

	migrationName := args[1]

	// Clean the migration name (replace spaces with underscores)
	migrationName = strings.ToLower(strings.ReplaceAll(migrationName, " ", "_"))

	// Generate timestamp
	timestamp := time.Now().Format("20060102150405")

	// Create migrations directory if it doesn't exist
	migrationsDir := "migrations"
	if err := os.MkdirAll(migrationsDir, 0755); err != nil {
		fmt.Printf("Error creating migrations directory: %v\n", err)
		os.Exit(1)
	}

	// Templates for the up and down migrations
	upTemplate := `-- Migration: %s (UP)
-- Created at: %s

BEGIN;

-- Add your migration SQL statements here

COMMIT;
`

	downTemplate := `-- Migration: %s (DOWN)
-- Created at: %s

BEGIN;

-- Add your rollback SQL statements here

COMMIT;
`

	// Generate up migration
	upFileName := fmt.Sprintf("%s_%s_up.sql", timestamp, migrationName)
	upFilePath := filepath.Join(migrationsDir, upFileName)

	if err := createMigrationFile(upFilePath, upTemplate, migrationName); err != nil {
		fmt.Printf("Error creating up migration file: %v\n", err)
		os.Exit(1)
	}

	// Generate down migration
	downFileName := fmt.Sprintf("%s_%s_down.sql", timestamp, migrationName)
	downFilePath := filepath.Join(migrationsDir, downFileName)

	if err := createMigrationFile(downFilePath, downTemplate, migrationName); err != nil {
		fmt.Printf("Error creating down migration file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully created migration files:")
	fmt.Printf("Up migration: %s\n", upFilePath)
	fmt.Printf("Down migration: %s\n", downFilePath)
}

func createMigrationFile(filePath, template, migrationName string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	content := fmt.Sprintf(template, migrationName, time.Now().Format("2006-01-02 15:04:05"))

	_, err = file.WriteString(content)
	return err
}
