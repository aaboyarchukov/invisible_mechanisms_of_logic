package codewithbags

import "math/rand/v2"

type RangeFloat struct {
	Min, Max float64
}

func (r RangeFloat) Random() float64 {
	return r.Min + rand.Float64()*(r.Max-r.Min)
}
