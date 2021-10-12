package greet

import (
	"fmt"
	"testing"
)

func ExampleGreet() {
	fmt.Println(Greet("James"))
	// Output:
	// James, hello my dear!
}

func TestGreet(t *testing.T) {
	s := Greet("Bond")
	if s != "Bond, hello my dear!" {
		t.Error("Expected", "Bond, hello my dear!", "Got: ", s)
	}
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
