package hoaretriples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		get  []int
		want []int
	}{
		{"empty array", []int{}, []int{}},
		{"one element", []int{1}, []int{1}},
		{"many elements 1", []int{3, 1, 2}, []int{1, 2, 3}},
		{"many elements 2", []int{40, 13, 4, 1}, []int{1, 4, 13, 40}},
		{"many elements 3", []int{7, 5, 6, 4, 3, 1, 2}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"many elements 4", []int{121, 40, 13, 4, 1}, []int{1, 4, 13, 40, 121}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			quickSort(tc.get, 0, len(tc.get)-1)

			assert.Equal(t, tc.want, tc.get,
				"FAIL: quickSort неверно отсортировал массив.\nОжидалось: %v\nПолучено:  %v",
				tc.want, tc.get,
			)
		})
	}
}
