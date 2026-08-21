package codewithbags

import "testing"

func TestAverageCalculator(t *testing.T) {
	calculator := NewAverageCalculator()

	cases := []struct {
		numbers  []int
		expected float64
	}{
		{[]int{1, 2, 3, 4, 5}, 3.0},
		{[]int{}, 0.0},
		{[]int{10, 20, 30}, 20.0},
	}

	for _, c := range cases {
		result := calculator.calculateAverage(c.numbers)

		if result != c.expected {
			t.Errorf("Expected %f, got %f", c.expected, result)
		}
	}
}
