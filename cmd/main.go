package main

import (
	"log"

	"github.com/umed-hotamov/url-shortener/config"
	"github.com/umed-hotamov/url-shortener/pkg/database"
	"github.com/umed-hotamov/url-shortener/pkg/logger"
)

func main() {
  cfg := config.GetConfig()
  logger, err := logger.NewLogger("info")
  if err != nil {
    log.Fatalf("Failed to create logger: %v", err)
  }

  db, err := database.InitDb(cfg, logger)
  if err != nil {
    logger.Fatal("Failed to create database")
  }
  defer db.Close()


}
