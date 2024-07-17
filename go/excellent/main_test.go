package main

import "testing"

func TestGreet(t *testing.T) {
	result := Greet("Test")
	expected := "Hello, Test!"
	if result != expected {
		t.Errorf("Greet(\"Test\") = %s; want %s", result, expected)
	}
}
