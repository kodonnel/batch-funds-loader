package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	result := hello()
	if result != "Hello World" {
		t.Errorf("failed expected %v got %v", "Hello World", result)
	}
}
