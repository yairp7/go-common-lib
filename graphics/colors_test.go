package graphics_test

import (
	"image/color"
	"testing"

	"github.com/yairp7/go-common-lib/graphics"
	"gopkg.in/stretchr/testify.v1/assert"
)

type testCase struct {
	Hex   string
	Color color.RGBA
	Err   error
}

var testCases = []testCase{
	{Hex: "#112233", Color: color.RGBA{17, 34, 51, 255}, Err: nil},
	{Hex: "#123", Color: color.RGBA{17, 34, 51, 255}, Err: nil},
	{Hex: "#000233", Color: color.RGBA{0, 2, 51, 255}, Err: nil},
	{Hex: "#023", Color: color.RGBA{0, 34, 51, 255}, Err: nil},
	{Hex: "invalid", Color: color.RGBA{0, 0, 0, 255}, Err: graphics.BadFormatErr},
	{Hex: "#abcd", Color: color.RGBA{0, 0, 0, 255}, Err: graphics.BadFormatErr},
	{Hex: "#-12", Color: color.RGBA{0, 0, 0, 255}, Err: graphics.BadFormatErr},
}

func Test_Colors_Hex2RGBA(t *testing.T) {
	for _, testCase := range testCases {
		clr, err := graphics.Hex2RGBA(testCase.Hex)
		assert.Equal(t, testCase.Err, err)
		assert.Equal(t, testCase.Color, clr)
	}
}
