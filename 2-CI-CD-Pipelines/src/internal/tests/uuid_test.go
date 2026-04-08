package test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/merteldem1r/DevOps-Fundamentals/2-CI-CD-Pipelines/src/utils"
)

func TestExample(t *testing.T) {
	id := utils.GenerateUUID()
	fmt.Println("Generated UUID:", id)

	if id == uuid.Nil {
		t.Error("generated UUID should not be Nil")
	}
}
