package main

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("perimeter 1", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Perimeter(rectangle)
		want := 40.0
		assertCorrectMessage(t, got, want)
	})
	t.Run("perimeter 2", func(t *testing.T) {
		rectangle := Rectangle{0.0, 0.0}
		got := Perimeter(rectangle)
		want := 0.0
		assertCorrectMessage(t, got, want)
	})

}

func TestArea(t *testing.T) {
	t.Run("area 1", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		got := Area(rectangle)
		want := 100.0
		assertCorrectMessage(t, got, want)
	})
	t.Run("area 2", func(t *testing.T) {
		rectangle := Rectangle{0.0, 0.0}
		got := Area(rectangle)
		want := 0.0
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, result, expected any) {
	t.Helper()
	if result != expected {
		t.Errorf("got %q want % q", result, expected)
	}
}
