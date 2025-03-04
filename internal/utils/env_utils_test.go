package utils

import "testing"

func TestCheckEnvNotNullOrDefault(t *testing.T) {
	got := CheckEnvNotNullOrDefault("myvalue", "default")
	want := "myvalue"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	gotStringNull := CheckEnvNotNullOrDefault("", "default")
	wantStringNull := "default"

	if gotStringNull != wantStringNull {
		t.Errorf("got %v, wanted %v", gotStringNull, wantStringNull)
	}

	gotInt := CheckEnvNotNullOrDefault("123", 0)
	wantInt := 123

	if gotInt != wantInt {
		t.Errorf("got %v, wanted %v", gotInt, wantInt)
	}

	gotIntNull := CheckEnvNotNullOrDefault("", 0)
	wantIntNull := 0

	if gotIntNull != wantIntNull {
		t.Errorf("got %v, wanted %v", gotIntNull, wantIntNull)
	}

	gotBool := CheckEnvNotNullOrDefault("true", false)
	wantBool := true

	if gotBool != wantBool {
		t.Errorf("got %v, wanted %v", gotBool, wantBool)
	}

	gotBoolNull := CheckEnvNotNullOrDefault("", false)
	wantBoolNull := false

	if gotBoolNull != wantBoolNull {
		t.Errorf("got %v, wanted %v", gotBoolNull, wantBoolNull)
	}
}
