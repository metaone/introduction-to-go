package simplestrings

import "testing"

const weekdays = "Monday Tuesday Wednesday Thursday Friday"

func TestIndex(t *testing.T) {
	var index int

	// test that an empty search string returns 0
	index = Index(weekdays, "")
	if index != 0 {
		t.Error("Expected that an empty search string returns 0, got ", index)
	}
}

func TestContains(t *testing.T) {
	var result bool

	// test that Tuesday is a weekday
	result = Contains(weekdays, "Tuesday")
	if !result {
		t.Error("Expected `Tuesday` is a weekday, got ", result)
	}

	// test that Sunday is not a weekday
	result = Contains(weekdays, "Sunday")
	if result {
		t.Error("Expected `Sunday` is not a weekday, got ", result)
	}

	// test that the string Monday is not found in the empty string
	result = Contains("", "Monday")
	if result {
		t.Error("Expected `Monday` is not found, got ", result)
	}
}

func TestHasPrefix(t *testing.T) {

}
