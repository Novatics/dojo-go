package main

import (
	"testing"
)

func testSaqueWithNegativeValueShouldThrowException(t *testing.T) {
	got :=  calculaSaque(100)
	want := 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}