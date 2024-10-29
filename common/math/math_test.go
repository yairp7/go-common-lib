package math_test

import (
	"testing"

	"github.com/yairp7/go-common-lib/common/math"
	"gopkg.in/stretchr/testify.v1/assert"
)

func Test_SliceMinMax(t *testing.T) {
	s := []int{6, 12, -56, 12, 5888, -2939}
	min, max := math.SliceMaxMin(s)
	assert.Equal(t, 5888, max)
	assert.Equal(t, -2939, min)
}
