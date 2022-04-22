package postgresql

import (
  "fmt"
  "log"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "github.com/RJHarvest/gin-api-template/config"
)

func getDBUrl(dbConfig *config.Config) string {
  return fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
    dbConfig.DBHost,
    dbConfig.DBUser,
    dbConfig.DBPassword,
    dbConfig.DBName,
    dbConfig.DBPort,
    dbConfig.DBSslmode,
  )
}

func InitDB() *gorm.DB {
  c := config.GetConfig()
  dsn := getDBUrl(c)
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatalf("[ InitDb ] Error: %v", err)
  }
  return db
}
