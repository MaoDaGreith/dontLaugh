//Package greet is just a simple test
package greet

import "fmt"

// Greet takes unlimited intereger values and returns the sum
func Greet(s string) string {
	return fmt.Sprintf("%v, hello my dear!", s)
}
