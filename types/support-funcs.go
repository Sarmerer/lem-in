// support-funcs.go
// Provides local mathematical functions.
// Package math only supports 64-bit floats, and the numbers that
// are being used in the graph package are integers.

package types

//Min returns the smaller of two integers.
func Min(x, y int) int {
	if x < y {
		return x
	} else if y < x {
		return y
	}
	return x
}

//Abs returns the absolute value of an integer.
func Abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}
