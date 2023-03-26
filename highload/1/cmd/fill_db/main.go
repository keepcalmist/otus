package main

import (
	"context"
	"fmt"
	"github.com/keepcalmist/otus/highload/1/internal/database"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	//host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable", port, user, password, dbName)

	log.Printf("trying to connect to db with dsn:%s", dsn)
	repo, err := database.NewRepo(ctx, dsn)
	if err != nil {
		log.Fatal("cant create repo:", err)
	}

	allUsers := parseRemoteCSV()

	for _, oneUser := range allUsers {
		err = repo.Users.InsertForCmd(ctx, "users", oneUser)
		if err != nil {
			log.Fatal(err)
		}
	}

}
