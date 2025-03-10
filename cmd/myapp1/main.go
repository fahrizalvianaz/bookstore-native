package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"simple-project-native/cmd/myapp1/config"
	"simple-project-native/cmd/myapp1/handler"
	"simple-project-native/cmd/myapp1/repository"
	"simple-project-native/cmd/myapp1/service"
	"simple-project-native/pkg/utils"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Config database failed", err)
	}
	db, err := utils.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Database connection failed", err)
	}

	defer db.Close()

	if err := createBooksTable(context.Background(), db); err != nil {
		log.Fatal("Failed to create table", err)
	}

	repoBook := repository.NewBookRepository(db)
	serviceBook := service.NewBookService(repoBook)
	handlerBook := handler.NewBookHandler(serviceBook)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /book", handlerBook.CreateBook)
	mux.HandleFunc("GET /book/{id}", handlerBook.GetBook)
	mux.HandleFunc("GET /books", handlerBook.GetAllBooks)

	repoUser := repository.NewUserRepository(db)
	serviceUser := service.NewUserService(repoUser)
	handlerUser := handler.NewUserHandler(serviceUser)

	mux.HandleFunc("POST /register", handlerUser.Register)
	mux.HandleFunc("POST /login", handlerUser.Login)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

}

func createBooksTable(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx,
		`CREATE TABLE IF NOT EXISTS books (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			author VARCHAR(255) NOT NULL 
			
		)
	`)
	return err
}
