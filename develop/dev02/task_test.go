package main

import (
	"testing"
)

func TestValidStrings(t *testing.T) {
	s := []struct {
		input  string
		expect string
	}{
		{"abcd", "abcd"},
		{"a2b2c2d2", "aabbccdd"},
		{"a0b0c0d0", ""},
		{"a3b4cd3", "aaabbbbcddd"},
	}

	for _, el := range s {
		result, err := unpack(el.input)
		if err != nil {
			t.Errorf("unexpected error:\n[%s] - string\n%s", el, err.Error())
		}
		if result != el.expect {
			t.Errorf("got %q, wanted %q", result, el.expect)
		}
	}
}

func TestInvalidStrings(t *testing.T) {
	s := []struct {
		input  string
		expect string
	}{
		{"34ab", "invalid string"},
		{`desa\`, "invalid string"},
	}

	for _, el := range s {
		result, err := unpack(el.input)
		if err == nil {
			t.Errorf("expected error, got %s", result)
		}
	}
}

func TestBackSlash(t *testing.T) {
	s := []struct {
		input  string
		expect string
	}{
		{`d\3w3`, "d3www"},
		{`\33`, "333"},
		{`a3\34`, "aaa3333"},
	}

	for _, el := range s {
		result, err := unpack(el.input)

		if err != nil {
			t.Errorf("unexpected error:\n[%s] - string\n%s", el, err.Error())
		}
		if result != el.expect {
			t.Errorf("got %q, wanted %q", result, el.expect)
		}
	}
}
