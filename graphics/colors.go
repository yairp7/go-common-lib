package graphics

import (
	"errors"
	"fmt"
	"image/color"
	"math"
)

var BadFormatErr = errors.New("bad format")

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

func Hex2RGBA(hex string) (clr color.RGBA, err error) {
	clr.A = 0xff
	defaultClr := clr

	if hex[0] != '#' {
		return clr, BadFormatErr
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = BadFormatErr
		return 0
	}

	switch len(hex) {
	case 7:
		clr.R = hexToByte(hex[1])<<4 + hexToByte(hex[2])
		clr.G = hexToByte(hex[3])<<4 + hexToByte(hex[4])
		clr.B = hexToByte(hex[5])<<4 + hexToByte(hex[6])
	case 4:
		clr.R = hexToByte(hex[1]) * 17
		clr.G = hexToByte(hex[2]) * 17
		clr.B = hexToByte(hex[3]) * 17
	default:
		err = BadFormatErr
	}

	if err != nil {
		clr = defaultClr
	}

	return
}
