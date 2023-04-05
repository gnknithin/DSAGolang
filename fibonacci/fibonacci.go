// Package fibonacci implements calculating fibonacci of given integer
// and return integer
package fibonacci

// Calculate returns its calculated Fibonacci number.
func Calculate(n int) int {
	// Calculate fibonacci of given integer
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Calculate(n-1) + Calculate(n-2)
	}
}
