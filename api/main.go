package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/aclel/sense/api/services"
	"github.com/aclel/sense/store"
	"golang.org/x/net/context"
)

func main() {
	dataSourceName := flag.String("db", "postgres://sense:sense@localhost/sense?sslmode=disable", "Database connection string")

	// Connect to the database
	db, err := store.NewDB(*dataSourceName)
	if err != nil {
		panic(err)
	}

	// Setup context which will be passed to all http handlers
	ctx := context.Background()
	context.WithValue(ctx, "db", db)
	router := services.NewRouter(ctx)

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
