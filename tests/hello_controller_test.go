package tests

import (
  "testing"
  "github.com/RJHarvest/gin-api-template/controllers/v1/hello"
)

func TestHelloController(t *testing.T) {
  message := hello.GetHello()
  if message != "Hello World" {
    t.Fatalf(`Expect: "Hello World", Got: %s instead`, message)
  }
}
