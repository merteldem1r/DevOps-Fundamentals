package test

import (
	"fmt"
	"testing"
)

func ExampleMessage() {
	fmt.Println("Testing package")
}

func TestExample(t *testing.T) {
	if 1+1 != 2 {
		t.Error("math is broken")
	}
}
