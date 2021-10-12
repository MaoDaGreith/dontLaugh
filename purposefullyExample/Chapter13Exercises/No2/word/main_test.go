package word

import (
	"fmt"
	"testing"
)

func TestUseCount(t *testing.T) {
	xs := UseCount("dos tres dos uno tres dos")
	for k, v := range xs {
		switch k {
		case "uno":
			if v != 1 {
				t.Error("Expected", 1, "Got", v)
			}
		case "dos":
			if v != 3 {
				t.Error("Expected", 3, "Got", v)
			}
		case "tres":
			if v != 2 {
				t.Error("Expected", 2, "Got", v)
			}
		}
	}
}

func TestCount(t *testing.T) {
	c := Count("dos tres dos uno tres dos cuatro cinco cuatro")
	if c != 9 {
		t.Error("Expected", 9, "Got", c)
	}
}

func ExampleCount() {
	fmt.Println(Count("One Two Two Three Four One"))
	// Output:
	// 6
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount("Juan Alejandro Antonio Rodrigo Ernesto Rikardo Juan Rikardo Alejandro")
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("uono duouo treie treie duo dui das dan dhog")
	}
}
