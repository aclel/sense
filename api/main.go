package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/aclel/sense/api/services"
	"github.com/aclel/sense/store"
)

func main() {
	dataSourceName := flag.String("db", "postgres://andrew:@localhost/sense?sslmode=disable", "Database connection string")

	db, err := store.NewDB(*dataSourceName)
	if err != nil {
		panic(err)
	}

	svc := services.NewService()
	router := services.NewRouter(db, svc)

	fmt.Println("Running...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
