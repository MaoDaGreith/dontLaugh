package dog

import (
	"fmt"
	"testing"
)

func TestYears(s *testing.T) {
	v := Years(7)

	if v != 49 {
		s.Error("Got", v, "Expected", 49)
	}
}

func TestYears2(t *testing.T) {
	x := YearsTwo(10)

	if x != 70 {
		t.Error("Got", x, "Expected", 70)
	}
}

func ExampleYears() {
	fmt.Println(Years(5))
	// Output:
	// 35
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(6))
	// Output:
	// 42
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(8)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(9)
	}
}
