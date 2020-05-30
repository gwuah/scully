package scully

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestScully(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		t.Error("No .env file found")
	}

	mapbox, err := New(os.Getenv("ACCESS_TOKEN"))

	if err != nil {
		t.Error(err)
	}

	response, err := mapbox.Matrix.GetMatrix(os.Getenv("TEST_POINTS"))

	if err != nil {
		t.Error(err)
	}

	log.Println(response)
}
