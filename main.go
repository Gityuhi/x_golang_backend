package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"x_golang_api/internal/handler"
	"x_golang_api/internal/repository"
	"x_golang_api/internal/router"
	"x_golang_api/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	url := fmt.Sprintf("postgresql://%s:%s@db:5432/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"))

	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	fmt.Println("success connected")
	defer pool.Close()

	userRepository := repository.NewUserRepository(pool)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	r := router.NewRouter(userHandler)

	log.Println("Server starting on port 8080...")

	r.Run()
}
