package config

import (
  "log"
  "os"
  "github.com/joho/godotenv"
)

type Config struct {
  Port       string
  DBHost     string
  DBUser     string
  DBPassword string
  DBName     string
  DBPort     string
  DBSslmode  string
}

func GetConfig() *Config {
  err := godotenv.Load(".env.local")
  if err != nil {
    log.Fatalf("[ setConfig ] Error loading env file: %v", err)
  }

  config := Config{
    Port: os.Getenv("PORT"),
    DBHost: os.Getenv("DB_HOST"),
    DBUser: os.Getenv("DB_USER"),
    DBPassword: os.Getenv("DB_PASSWORD"),
    DBName: os.Getenv("DB_NAME"),
    DBPort: os.Getenv("DB_PORT"),
    DBSslmode: os.Getenv("DB_SSLMODE"),
  }
  return &config
}
