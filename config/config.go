package config

import (
  "log"
  "github.com/spf13/viper"
)

func Config() {
  viper.SetConfigName(".env")

}