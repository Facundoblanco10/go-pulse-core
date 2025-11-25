package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Facundoblanco10/go-pulse-core/internal/api"
	"github.com/Facundoblanco10/go-pulse-core/internal/jobs"
	"github.com/Facundoblanco10/go-pulse-core/internal/storage"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	env := os.Getenv("ENVIRONMENT")
	if env == "" || env == "development" {
		if err := godotenv.Load(); err != nil {
			log.Println("⚠️ .env file not found, relying on system env")
		}
	}
}

func main() {
	env := getEnv("ENVIRONMENT", "development")
	dsn := buildDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to initialize database, got error %v", err)
	}

	if env != "production" {
		if err := db.AutoMigrate(&storage.JobModel{}); err != nil {
			log.Fatalf("failed to migrate: %v", err)
		}
	}

	jobRepo := storage.NewJobRepository(db)
	jobSvc := jobs.NewService(jobRepo)
	r := api.NewRouter(jobSvc)

	port := getEnv("PORT", "8080")
	log.Printf("PulseCore running in %s mode on port %s\n", env, port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}

func buildDSN() string {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	pass := getEnv("DB_PASSWORD", "postgres")
	name := getEnv("DB_NAME", "pulsecore")

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, pass, name, port,
	)
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
