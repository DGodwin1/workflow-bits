package main

import "testing"

func TestGenerateHello(t *testing.T) {
	got := GenerateHello("david")
	want := "hello david"
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestGenerateFood(t *testing.T) {
	got := GenerateFood([]string{"beef"})
	want := "beef+rice"
	if got != want {
		t.Errorf("whoops!")
	}
}
