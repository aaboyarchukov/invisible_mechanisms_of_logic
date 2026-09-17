package hoaretriples

import "math"

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

var (
	InvalidMax = math.MinInt
	MinimumInt = math.MinInt
)

// P: {len(arr) > 0}
// C: findMax(arr)
// Q: {result = max(arr)}
// I: {result = max(arr[0:i+1])}
func findMax(arr []int) int {
	result := MinimumInt

	if len(arr) == 0 {
		return InvalidMax
	}

	for number := range arr {
		if number > result {
			result = number
		}
	}

	return result
}

// P: {len(arr) > 0}
// C: chunkArray(arr, left, right)
// Q: {arr[0] < arr[1] < ... < arr[n]}}
// I: {left <= right, arr[left] <= arr[left+1] <= ... <= arr[right]}
func chunkArray(arr []int, left, right int) int {
	middle := (left + right) / 2

	target := arr[middle]

	for left <= right {
		for arr[left] < target {
			left++
		}

		for arr[right] > target {
			right--
		}

		if left >= right {
			return right
		}

		arr[left], arr[right] = arr[right], arr[left]

	}

	return right
}

// P: {len(arr) > 0}
// C: quickSort(arr, left, right)
// Q: {arr[0] < arr[1] < ... < arr[n]}}
// I: {left <= right}
// I1: {arr[left] <= arr[left+1] <= ... <= arr[partitionIndx]}
// I2: {arr[partitionIndx + 1] <= arr[partitionIndx+2] <= ... <= arr[right]}
func quickSort(arr []int, left, right int) {
	if len(arr) < 2 {
		return
	}

	if left < right {
		partitionIndx := chunkArray(arr, left, right)
		quickSort(arr, left, partitionIndx)
		quickSort(arr, partitionIndx+1, right)
	}

}
