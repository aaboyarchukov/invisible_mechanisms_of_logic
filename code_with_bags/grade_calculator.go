package codewithbags

import "errors"

type GradeCalculator struct{}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{}
}

var (
	ErrNegativeGrade  = errors.New("grade is negative")
	ErrZeroSumOfGrade = errors.New("sum of grades is zero")
)

var (
	ZeroAvgGrade    = 0.0
	InvalidAvgGrade = -1.0
)

// conds:
// - zero len grades
// - negative grades
// - sum of grades equal to zero
// - grades are numbers not runes (it's check by compiler)
func (gc *GradeCalculator) calculateAverage(grades []int) (float64, error) {
	countGrades := len(grades)

	if countGrades == 0 {
		return ZeroAvgGrade, nil
	}

	sum := 0
	for _, grade := range grades {
		if grade < 0 {
			return InvalidAvgGrade, ErrNegativeGrade
		}

		sum += grade
	}

	if sum == 0 {
		return InvalidAvgGrade, ErrZeroSumOfGrade
	}

	return float64(sum) / float64(countGrades), nil
}
