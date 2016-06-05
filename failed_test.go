package main

import (
	"testing"
)

func TestHelloWithoutForcedFail(t *testing.T) {
	greeting := hello()
	expected := "Hello, World!" // Intentional mismatch
	if greeting != expected {
		t.Errorf("greeting=\"%s\", but we expected \"%s\"",
			greeting, expected)
	}
}
