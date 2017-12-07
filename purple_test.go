package main

import "testing"

func TestGetColor(t *testing.T) {
	expected := "purple"
	actual := GetColor()

	if actual != expected {
		t.Errorf("getColor() == %q, expected %q", actual, expected)
	}
}
