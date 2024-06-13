package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello Go"

	got := "Hello Go"

	if want != got {
		t.Fatalf("want %s, got %s\n", want, got)
	}
}
