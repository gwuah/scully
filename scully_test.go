package scully

import (
	"testing"
)

func TestScully(t *testing.T) {
	accessToken := "3974398477387473"
	mapbox, err := New(accessToken)

	if err != nil {
		t.Error(t)
	}

	if mapbox.Greet() != "Hello World" {
		t.Error("Greeting doesn't match")
	}
}
