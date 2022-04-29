package main

import (
	"reflect"
	"testing"
)

func TestCut(t *testing.T) {
	flgs := customFlags{
		fieldsFlag:    []int{2, 3},
		delimiterFlag: " ",
		separatedFlag: true,
	}

	cut := customCut{
		strs: []string{
			"ab bc cd",
			"er 23 43",
			"re er re",
		},
		flags: flgs,
	}

	expect := []string{"bc cd", "23 43", "er re"}
	result := cut.Cut()

	if !reflect.DeepEqual(expect, result) {
		t.Errorf("wanted %q, got %q", expect, result)
	}
}
