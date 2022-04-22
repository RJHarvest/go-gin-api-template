package postgresql

import (
  "fmt"
  "log"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

type DBConfig struct {
  Host     string
  User     string
  Password string
  DBName   string
  Port     int
  Sslmode  string
}

func buildDBConfig() *DBConfig {
  dbConfig := DBConfig{
    Host: "app-26de4dde-2df4-4e19-8164-f26387fcbcda-do-user-9709129-0.b.db.ondigitalocean.com",
    User: "deskchaser-db",
    Password: "y0CqiWU32q0lgSUu",
    DBName: "deskchaser-db",
    Port: 25060,
    Sslmode: "require",
  }
  return &dbConfig
}

func getDBUrl(dbConfig *DBConfig) string {
  return fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
    dbConfig.Host,
    dbConfig.User,
    dbConfig.Password,
    dbConfig.DBName,
    dbConfig.Port,
    dbConfig.Sslmode,
  )
}

func InitDB() (*gorm.DB) {
  dbConfig := buildDBConfig()
  dsn := getDBUrl(dbConfig)
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatalf("[ InitDb ] Error: %v", err)
  }
  return db
}
