package codewithbags

type AverageCalculator struct {
}

func NewAverageCalculator() *AverageCalculator {
	return &AverageCalculator{}
}

func (c *AverageCalculator) calculateAverage(numbers []int) float64 {
	countNumbers := len(numbers)
	if countNumbers == 0 {
		return 0
	}

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	return float64(sum) / float64(countNumbers)
}
