package hoaretriples

var (
	LESS    = -1
	EQUAL   = 0
	GREATER = 1
)

// P: a = n, b = m, compareResult = NONE
// C: max(a, b)
// Q: compareResult {GREATER, LESS, EQUAL}
func max(a, b int) int {
	if a > b {
		return GREATER
	}

	if a < b {
		return LESS
	}

	return EQUAL
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
