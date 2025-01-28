package main

import (
	"fmt"
	"os"

	"github.com/byvinesse/vinance-db-migrations/cmd/generate"
	"github.com/byvinesse/vinance-db-migrations/cmd/migrate"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "generate":
		generate.Generate()
	case "migrate":
		migrate.Migrate()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: go run main.go <command> [arguments]")
	fmt.Println("\nAvailable commands:")
	fmt.Println("  migrate  - Run database migrations")
	fmt.Println("  generate - Generate new migration files")
	fmt.Println("\nExamples:")
	fmt.Println("  go run main.go generate create_users_table")
	fmt.Println("  go run main.go migrate")
}
