package test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestExample(t *testing.T) {
	id := uuid.New()
	fmt.Println("Generated UUID:", id)

	if id == uuid.Nil {
		t.Error("generated UUID should not be Nil")
	}
}
