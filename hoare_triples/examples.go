package hoaretriples

// P: a = n, b = m, compareResult = True
// C: max(a, b)
// Q: (compareResult = a || compareResult = b ) && (compareResult >= a && compareResult >= b)
func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// P: x = n
// C: abs(a)
// Q: x {-n, n}
func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// P: a = n, b = m, compareResult = NONE
// C: MaxOfAbs(a, b)
// Q: compareResult {GREATER, LESS, EQUAL}
func MaxOfAbs(a, b int) int {
	return max(abs(a), abs(b))
}
