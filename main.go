package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func create() {
	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s create <firstname> <lastname>\n", os.Args[0])
		os.Exit(1)
	}
	err := CreatePerson(context.Background(), PersonParams{
		FirstName: os.Args[2],
		LastName:  os.Args[3],
	})
	if err != nil {
		log.Fatalf("Creating person error: %v", err)
	}
}

func update() {
	if len(os.Args) < 5 {
		fmt.Fprintf(os.Stderr, "Usage: %s update <id> <firstname> <lastname>\n", os.Args[0])
		os.Exit(1)
	}
	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid id, must be a integer\n")
		fmt.Fprintf(os.Stderr, "Usage: %s update <id> <firstname> <lastname>\n", os.Args[0])
		os.Exit(1)
	}

	err = UpdatePerson(context.Background(), id, PersonParams{
		FirstName: os.Args[3],
		LastName:  os.Args[4],
	})
	if err != nil {
		log.Fatalf("Creating person error: %v", err)
	}
}

func list() {
	people, err := ListPeople(context.Background())
	if err != nil {
		log.Fatalf("Listing people error: %v", err)
	}
	for _, person := range people {
		fmt.Printf("[%d] %s %s\n", person.ID, person.FirstName, person.LastName)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
		os.Exit(1)
	}

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading local env files")
	}

	// Setup Database
	db, err = sql.Open("postgres", dbURI())
	if err != nil {
		log.Fatalf("database connection error: %v", err)
	}

	cmd := os.Args[1]
	switch cmd {
	case "list":
		list()
	case "create":
		create()
	case "update":
		update()
	default:
		fmt.Fprintf(os.Stderr, "Invalid command\n")
		fmt.Fprintf(os.Stderr, "Usage: %s <cmd> ...\n", os.Args[0])
		os.Exit(1)
	}
}

func dbURI() string {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)
}

const usage = `Usage: %s <command> ...

Commands:

    list

    create <firstname> <lastname>

    update <id> <firstname> <lastname>

`
