package hoaretriples

// P: a = n, b = m, result = True
// C: max(a, b)
// Q: (result = a || result = b ) && (result >= a && result >= b)
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// P: true
// C: abs(x)
// Q: result = |x|
func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// P: True
// C: MaxOfAbs(a, b)
// Q: max(|a|, |b|)
func MaxOfAbs(a, b int) int {
	return max(abs(a), abs(b))
}
