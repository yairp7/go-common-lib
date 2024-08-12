package graphics

import (
	"fmt"
	"image/color"
	"math"
)

func ColorDistance(c1 color.Color, c2 color.Color) float64 {
	r1, g1, b1, _ := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()
	dR := math.Abs(float64(r1-r2)) / 255
	dG := math.Abs(float64(g1-g2)) / 255
	dB := math.Abs(float64(b1-b2)) / 255
	return (dR + dG + dB) / 3 * 100
}

func RGB2Hex(clr color.Color) string {
	rgba := color.RGBAModel.Convert(clr).(color.RGBA)
	return fmt.Sprintf("#%.2x%.2x%.2x", rgba.R, rgba.G, rgba.B)
}
