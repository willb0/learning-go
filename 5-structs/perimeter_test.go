package main

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		assertCorrectMessage(t, got, want)
	}
	t.Run("perimeter 1", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkPerimeter(t, rectangle, 40.0)
	})
	t.Run("perimeter 2", func(t *testing.T) {
		rectangle := Rectangle{0.0, 0.0}
		checkPerimeter(t, rectangle, 0.0)
	})
	t.Run("perimeter 3", func(t *testing.T) {
		circle := Circle{10.0}
		checkPerimeter(t, circle, 10*math.Pi*2)
	})

}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		assertCorrectMessage(t, got, want)
	}
	t.Run("area rectangle 1", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})
	t.Run("area rectangle 2", func(t *testing.T) {
		rectangle := Rectangle{0.0, 0.0}
		checkArea(t, rectangle, 0.0)
	})
	t.Run("area circle 1", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}

func assertCorrectMessage(t testing.TB, result, expected any) {
	t.Helper()
	if result != expected {
		t.Errorf("got %g want %g", result, expected)
	}
}
