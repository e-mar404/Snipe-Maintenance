package snipe

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestEnvFile (t *testing.T) {
  err := godotenv.Load("../../.env")

  if err != nil {
    t.Fatal("there is no env file at the root of the proj")
  }
}

func TestEnvKey (t *testing.T) {
  godotenv.Load("../../.env")

  API_KEY := os.Getenv("API_KEY")
  exists := len(API_KEY) > 0

  if !exists {
    t.Fatal("no api key in env file")
  }
}
