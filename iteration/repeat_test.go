package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}

}

func BenchmarkRepeat(b *testing.B) {

	for i := 0; i < b.N; i++ {

		Repeat("a", 10)
	}
}

func ExampleRepeat() string {

	repeated := Repeat("a", 20)
	return repeated

	// This function should repeat a given string n times
}

func TestJoin(t *testing.T) {
	array := []string{"C", "a", "r", "l", "o", ""}
	got := Join(array)
	want := "Carlos"

	fmt.Println("Brabo")
	if got != want {

		t.Errorf("Got: %s Want: %s", got, want)
	}
}
