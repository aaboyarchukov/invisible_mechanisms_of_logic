package codewithbags

import (
	"errors"
	"math"
)

type GradeCalculator struct{}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{}
}

var (
	ErrNegativeGrade = errors.New("grade is negative")
	ErrEmptyGrades   = errors.New("grades are empty")
	ErrOverflow      = errors.New("int overflow")
	ErrGradeTooLarge = errors.New("grade is too large")
)

var (
	ZeroAvgGrade    = 0.0
	InvalidAvgGrade = -1.0
	MaxGrade        = 100
)

// conds:
// - zero len grades
// - negative grades
// - grades are numbers not runes (it's check by compiler)
func (gc *GradeCalculator) calculateAverage(grades []int) (float64, error) {
	countGrades := len(grades)

	if countGrades == 0 {
		return InvalidAvgGrade, ErrEmptyGrades
	}

	if countGrades > math.MaxInt/MaxGrade {
		return InvalidAvgGrade, ErrOverflow
	}

	sum := 0
	for _, grade := range grades {
		if grade < 0 {
			return InvalidAvgGrade, ErrNegativeGrade
		}

		if grade > MaxGrade {
			return InvalidAvgGrade, ErrGradeTooLarge
		}

		sum += grade
	}

	return float64(sum) / float64(countGrades), nil
}
