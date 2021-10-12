//Package exampl is just a simple test
package exampl

// Muscles takes unlimited intereger values and returns the sum
func Muscles(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}
