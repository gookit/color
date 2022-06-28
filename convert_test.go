package color

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRgb2basic(t *testing.T) {
	assert.Equal(t, uint8(31), Rgb2basic(134, 56, 56, false))
	assert.Equal(t, uint8(41), Rgb2basic(134, 56, 56, true))
	assert.Equal(t, uint8(46), Rgb2basic(57, 187, 226, true))
}

func TestHex2basic(t *testing.T) {
	assert.Equal(t, uint8(95), Hex2basic("fd7cfc"))
	assert.Equal(t, uint8(105), Hex2basic("fd7cfc", true))
}

func TestHslToRgb(t *testing.T) {
	// red #ff0000	255,0,0  0,100%,50%
	rgbVal := HslToRgb(0, 1, 0.5)
	// fmt.Println(rgbVal)
	assert.Equal(t, []uint8{255, 0, 0}, rgbVal)

	rgbVal = HslIntToRgb(0, 100, 50)
	// fmt.Println(rgbVal)
	assert.Equal(t, []uint8{255, 0, 0}, rgbVal)

	rgbVal = HslIntToRgb(0, 100, 25)
	// fmt.Println(rgbVal)
	assert.Equal(t, []uint8{128, 0, 0}, rgbVal)

	// darkgray	 #a9a9a9 169,169,169 0,0%,66%
	rgbVal = HslIntToRgb(0, 0, 66)
	fmt.Println(rgbVal)
	assert.Equal(t, []uint8{168, 168, 168}, rgbVal)

	rgbVal = HslToRgb(0, 0, 0.6627)
	fmt.Println(rgbVal)
	assert.Equal(t, []uint8{169, 169, 169}, rgbVal)

	hslVal := RgbToHslInt(rgbVal[0], rgbVal[1], rgbVal[2])
	fmt.Println(hslVal)
	assert.Equal(t, []int{0, 0, 66}, hslVal)

	hslFVal := RgbToHsl(rgbVal[0], rgbVal[1], rgbVal[2])
	fmt.Println("rgb:", rgbVal, "=> hsl:", hslFVal)
	assert.NotEmpty(t, hslFVal)

	hslFVal = RgbToHsl(57, 187, 226)
	fmt.Println("rgb: 57, 187, 226 => hsl:", hslFVal)
	assert.NotEmpty(t, hslFVal)
}
