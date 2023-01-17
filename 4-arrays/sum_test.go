package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertCorrectMessage(t, got, want)
	})

	t.Run("collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6
		assertCorrectMessage(t, got, want)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{2, 9})
	want := []int{3, 11}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want % q", got, want)
	}
}

func assertCorrectMessage(t testing.TB, result, expected any) {
	t.Helper()
	if result != expected {
		t.Errorf("got %q want % q", result, expected)
	}
}
