package hello

import (
  "log"
)

func GetHello() string {
  log.Println("[ GetHello ] Start method")
  return "Hello World"
}
