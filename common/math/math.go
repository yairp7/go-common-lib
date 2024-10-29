package math

import "math"

type Number interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | float32 | float64
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func ToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

func SliceMaxMin[T Number](s []T) (min T, max T) {
	if len(s) == 0 {
		return min, max
	}

	min, max = s[0], s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}

		if s[i] < min {
			min = s[i]
		}
	}

	return min, max
}
