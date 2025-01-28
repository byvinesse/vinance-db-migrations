package migrate

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Migration struct {
	Version   string
	Filename  string
	Applied   bool
	AppliedAt *time.Time
}

func init() {
	// Load DB configs
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}
}

func Migrate() {
	host := getEnvOrDefault("DB_HOST", "localhost")
	port := getEnvOrDefault("DB_PORT", "5432")
	user := getEnvOrDefault("DB_USER", "postgres")
	pass := getEnvOrDefault("DB_PASS", "postgres")
	name := getEnvOrDefault("DB_NAME", "vinance")

	log.Printf("Connecting to database %s on %s:%s\n", name, host, port)

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	defer db.Close()

	// Test connection
	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	log.Println("Successfully connected to database")

	// Ensure migrations table exists
	err = createMigrationsTable(db)
	if err != nil {
		log.Fatal("Error creating migrations table:", err)
	}

	// Get all migration files
	migrations, err := getMigrationFiles("migrations")
	if err != nil {
		log.Fatal("Error reading migration files:", err)
	}

	// Get applied migrations from database
	appliedMigrations, err := getAppliedMigrations(db)
	if err != nil {
		log.Fatal("Error getting applied migrations:", err)
	}

	// Apply pending migrations
	for _, migration := range migrations {
		if !isMigrationApplied(migration.Version, appliedMigrations) {
			fmt.Printf("Applying migration: ", migration.Filename)
			err := applyMigration(db, migration)
			if err != nil {
				log.Fatal("Error applying migration:", err)
			}
		}
	}
}

func getEnvOrDefault(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func createMigrationsTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version VARCHAR(255) PRIMARY KEY,
			applied_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);
	`

	_, err := db.Exec(query)
	return err
}

func getMigrationFiles(dir string) ([]Migration, error) {
	var migrations []Migration

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".sql") {
			version := strings.Split(file.Name(), "_")[0]
			migrations = append(migrations, Migration{
				Version:  version,
				Filename: filepath.Join(dir, file.Name()),
			})
		}
	}

	// Sort migrations by version
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})

	return migrations, nil
}

func getAppliedMigrations(db *sql.DB) (map[string]time.Time, error) {
	applied := make(map[string]time.Time)

	rows, err := db.Query("SELECT version, applied_at FROM schema_migrations")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var version string
		var appliedAt time.Time
		if err := rows.Scan(&version, &appliedAt); err != nil {
			return nil, err
		}
		applied[version] = appliedAt
	}

	return applied, rows.Err()
}

func isMigrationApplied(version string, applied map[string]time.Time) bool {
	_, exists := applied[version]
	return exists
}

func applyMigration(db *sql.DB, migration Migration) error {
	content, err := ioutil.ReadFile(migration.Filename)
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Execute the migration script
	_, err = tx.Exec(string(content))
	if err != nil {
		tx.Rollback()
		return err
	}

	// Record the migration
	_, err = tx.Exec("INSERT INTO schema_migrations (version) VALUES($1)", migration.Version)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
