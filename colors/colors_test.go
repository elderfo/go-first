package colors

import "testing"

func AssertColorMatch(t *testing.T, actual string, expected string, err error) {
	if err != nil {
		t.Errorf("GetColor() err = %q, expected nil", err)
		return
	}

	if actual != expected {
		t.Errorf("getColor() == %q, expected %q", actual, expected)
		return
	}
}

func TestGetColorWithPurple(t *testing.T) {
	expected := "purple"
	actual, err := GetColor(Purple)

	AssertColorMatch(t, actual, expected, err)
}

func TestGetColorWithRed(t *testing.T) {
	expected := "red"
	actual, err := GetColor(Red)

	AssertColorMatch(t, actual, expected, err)
}

func TestGetColorWithBlue(t *testing.T) {
	expected := "blue"
	actual, err := GetColor(Blue)

	AssertColorMatch(t, actual, expected, err)
}

func TestGetColorWithUnknownArg(t *testing.T) {
	actual, err := GetColor(9)

	if actual != "" {
		t.Errorf("GetColor(9) returned color %q", actual)
		return
	}

	if err == nil {
		t.Errorf("Expected err to have a value")
		return
	}
}
