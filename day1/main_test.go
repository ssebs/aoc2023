package main

import "testing"

func TestGetDigits(t *testing.T) {
	samples := map[string]int{"2five524": 24, "fx3": 33, "363mtk": 33}

	for s, want := range samples {
		got := GetDigits(s)
		if got != want {
			t.Fatalf("got %d want %d", got, want)
		}
	}
}
