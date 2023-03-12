package main

import (
	"context"
	"fmt"
	"github.com/keepcalmist/otus/highload/1/internal/database"
	"github.com/keepcalmist/otus/highload/1/internal/delivery/http"
	"log"
	defaulHttp "net/http"
	"os"
)

func main() {
	ctx := context.Background()
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	//dsn := "postgresql://postgres:root@postgres/postgres?sslmode=disable"
	log.Printf("trying to connect to db with dsn:%s", dsn)
	repo, err := database.NewRepo(ctx, dsn)
	if err != nil {
		log.Fatal("cant create repo:", err)
	}

	router := http.NewRouter(repo)

	if err := defaulHttp.ListenAndServe("localhost:8080", router); err != nil {
		log.Fatal(err)
		return
	}
}
