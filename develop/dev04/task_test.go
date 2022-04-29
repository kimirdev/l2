package main

import (
	"reflect"
	"testing"
)

func TestStrings(t *testing.T) {
	strs := []string{
		"пятак",
		"пятка",
		"тяпка",
		"листок",
		"слиток",
		"выфв",
		"ыввф",
		"столик",
	}
	expect := map[string][]string{
		"выфв":   {"выфв", "ыввф"},
		"листок": {"листок", "слиток", "столик"},
		"пятак":  {"пятак", "пятка", "тяпка"},
	}

	result := getAnagrams(strs)

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("got %s, wanted %s", result, expect)
	}
}
