package main

import (
	"reflect"
	"testing"
)

func TestStrings(t *testing.T) {
	s := []struct {
		input  []string
		expect []string
	}{
		{[]string{"cde", "bca", "abc", "zxc"}, []string{"abc", "bca", "cde", "zxc"}},
		{[]string{"2", "1", "4", "4", "5"}, []string{"1", "2", "4", "4", "5"}},
	}

	for _, el := range s {
		res := sortUtil(el.input, 0, false, false, false)
		if eq := reflect.DeepEqual(res, el.expect); !eq {
			t.Errorf("got %s, wanted %s", el.input, el.expect)
		}
	}
}

func TestUniq(t *testing.T) {
	s := []struct {
		input  []string
		expect []string
	}{
		{[]string{"cde", "bca", "abc", "zxc", "abc"}, []string{"abc", "bca", "cde", "zxc"}},
		{[]string{"2", "1", "4", "4", "5"}, []string{"1", "2", "4", "5"}},
	}

	for _, el := range s {
		res := sortUtil(el.input, 0, false, false, true)
		if eq := reflect.DeepEqual(res, el.expect); !eq {
			t.Errorf("got %s, wanted %s", el.input, el.expect)
		}
	}
}

func TestReverse(t *testing.T) {
	s := []struct {
		input  []string
		expect []string
	}{
		{[]string{"cde", "bca", "abc", "zxc"}, []string{"zxc", "cde", "bca", "abc"}},
		{[]string{"2", "1", "4", "4", "5"}, []string{"5", "4", "4", "2", "1"}},
	}

	for _, el := range s {
		res := sortUtil(el.input, 0, false, true, false)
		if eq := reflect.DeepEqual(res, el.expect); !eq {
			t.Errorf("got %s, wanted %s", el.input, el.expect)
		}
	}
}
