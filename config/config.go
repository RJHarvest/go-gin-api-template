package config

import (
  "log"
  "os"
  "github.com/joho/godotenv"
  "github.com/RJHarvest/gin-api-template/utils"
)

type Config struct {
  Port              string
  DBHost            string
  DBUser            string
  DBPassword        string
  DBName            string
  DBPort            string
  DBSslmode         string
  CorsAllowOrigins  []string
}

// Constructor function
func (config *Config) fillDefaults() {
  if config.Port == "" {
    config.Port = "3000"
  }
  if len(config.CorsAllowOrigins) == 0 {
    config.CorsAllowOrigins = []string{"*"}
  }
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
    CorsAllowOrigins: utils.ParseStringToArray(os.Getenv("CORS_ALLOW_ORIGINS")),
  }

  config.fillDefaults()
  log.Printf("[ Config ] %v", &config)

  return &config
}
