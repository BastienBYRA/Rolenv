package config

import (
	"slices"
	"testing"
)

func TestParseKeyValuePairs(t *testing.T) {
	got := ParseKeyValuePairs("mykey:myvalue;mykey2:myvalue2")
	want := []string{"mykey:myvalue", "mykey2:myvalue2"}

	if !slices.Equal(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	gotEmpty := ParseKeyValuePairs("")
	wantEmpty := []string{}

	if !slices.Equal(gotEmpty, wantEmpty) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestParseBoolEnv(t *testing.T) {
	boolTrue := []string{"true", "yes", "True", "TRUE"}
	boolFalse := []string{"false", "no", "False", "FALSE", ""}
	for _, value := range boolTrue {
		got := parseBoolEnv(value)
		want := true

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

	for _, value := range boolFalse {
		got := parseBoolEnv(value)
		want := false

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}

func TestParseList(t *testing.T) {
	got := parseList("apple;banana;cherry")
	want := []string{"apple", "banana", "cherry"}

	if !slices.Equal(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	gotEmpty := parseList("")
	wantEmpty := []string{}

	if !slices.Equal(gotEmpty, wantEmpty) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestParsePositiveNumber(t *testing.T) {
	got := parsePositiveNumber("10")
	want := 10

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	gotZero := parsePositiveNumber("0")
	wantZero := 0

	if gotZero != wantZero {
		t.Errorf("got %v, wanted %v", gotZero, wantZero)
	}

	gotNegative := parsePositiveNumber("-1")
	wantNegative := -1

	if gotNegative != wantNegative {
		t.Errorf("got %v, wanted %v", gotNegative, wantNegative)
	}
}
