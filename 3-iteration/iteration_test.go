package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat_1(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q got %q", expected, repeated)
	}
}

func TestRepeat_2(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", i)
	}
}

func ExampleRepeat() {
	res := Repeat("a", 5)
	fmt.Println(res)
	// "aaaaa"
}
