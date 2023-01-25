package utils

import (
	"strings"
)

func ParseStringToArray(stringToParse string) []string {
  if stringToParse == "" {
    return []string{}
  }
  return strings.Split(stringToParse, ",")
}