package codewithbags

import (
	"encoding/binary"
	"errors"
	"testing"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/require"
)

// 1. handle test with viewing through the eyes

// 2. simple auto test
func TestGradeCalculator_Simple(t *testing.T) {
	t.Helper()

	gradeCalculator := NewGradeCalculator()

	cases := []struct {
		grades   []int
		expected float64
		err      error
	}{
		{[]int{1, 2, 3, 4, 5}, 3.0, nil},                                     // happy-path case
		{[]int{}, InvalidAvgGrade, ErrEmptyGrades},                           // no grades
		{[]int{10, 20, -30}, InvalidAvgGrade, ErrNegativeGrade},              // negative grades
		{[]int{-10, -20, -30}, InvalidAvgGrade, ErrNegativeGrade},            // negative sum
		{[]int{1000, 12222, 23123123123}, InvalidAvgGrade, ErrGradeTooLarge}, // too large grade
	}

	for _, c := range cases {
		result, err := gradeCalculator.calculateAverage(c.grades)

		require.Equalf(t, c.err, err, "Unexpected err: %w, expected = %w", err, c.err)
		require.Equalf(t, c.expected, result, "Err with calculate: expected = %f, result = %f", c.expected, result)

	}
}

func intsToBytes(nums []int) []byte {
	buf := make([]byte, 0, len(nums)*2)
	for _, n := range nums {
		buf = binary.AppendVarint(buf, int64(n))
	}
	return buf
}

func bytesToInts(data []byte) []int {
	nums := make([]int, 0, len(data))
	for len(data) > 0 {
		v, n := binary.Varint(data)
		if n <= 0 {
			break
		}

		if int64(int(v)) != v {
			break
		}
		nums = append(nums, int(v))
		data = data[n:]
	}
	return nums
}

var matchErrorsMap = map[error]bool{
	ErrEmptyGrades:   true,
	ErrGradeTooLarge: true,
	ErrInvalidValue:  true,
	ErrNegativeGrade: true,
}

func matchErrors(targetErr error) bool {
	return matchErrorsMap[targetErr]
}

// 3. fazzing test
func FuzzGradeCalculator(t *testing.F) {
	t.Helper()
	gradeCalculator := NewGradeCalculator()

	var (
		fakeDataBufferSize int = 100
		arraySize          int = 10

		minRange = 0
		maxRange = 100
	)

	fake := faker.New()
	for range fakeDataBufferSize {
		templeArray := make([]int, 0, arraySize)
		for range arraySize {
			templeArray = append(templeArray, int((float64(fake.IntBetween(minRange, maxRange)))))
		}

		numbers := intsToBytes(templeArray)
		t.Add(numbers)
	}

	t.Fuzz(func(t *testing.T, grades []byte) {
		numbers := bytesToInts(grades)
		result, err := gradeCalculator.calculateAverage(numbers)

		if err != nil && !matchErrors(err) {
			t.Fatalf("Unexpected err: %v", err)
		}

		if err != nil {
			require.Equal(t, result, -1.0)

			return
		}

		require.GreaterOrEqual(t, result, 0.0, "Result is negative is wrong")

	})
}

// 4. handle test with handing run programm
func TestRunGradeCalculator(t *testing.T) {
	gradeCalculator := NewGradeCalculator()

	result, err := gradeCalculator.calculateAverage([]int{1, 2, 3})

	if err != nil {
		t.Fatalf("didn't expect a mistake, but err is %v", err)
	}

	if result != 2 {
		t.Fatalf("result invalid = %f", result)
	}

	result, err = gradeCalculator.calculateAverage([]int{0})

	if err != nil {
		t.Fatalf("didn't expect a mistake, but err is %v", err)
	}

	if result != 0 {
		t.Fatalf("result invalid = %f", result)
	}

	result, err = gradeCalculator.calculateAverage([]int{})

	if !errors.Is(err, ErrEmptyGrades) {
		t.Fatalf("didn't expect a mistake, but err is %v", err)
	}

	if result != InvalidAvgGrade {
		t.Fatalf("result invalid = %f", result)
	}

	result, err = gradeCalculator.calculateAverage([]int{-1, -2, -3})

	if !errors.Is(err, ErrNegativeGrade) {
		t.Fatalf("didn't expect a mistake, but err is %v", err)
	}

	if result != InvalidAvgGrade {
		t.Fatalf("result invalid = %f", result)
	}
}
