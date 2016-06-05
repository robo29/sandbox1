package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	greeting := hello()
	expected := "Hello, World!"
	if greeting != expected {
		t.Errorf("greeting=\"%s\", but we expected \"%s\"",
			greeting, expected)
	}
}
