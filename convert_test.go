package color

import (
	"fmt"
	"testing"

	"github.com/gookit/assert"
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
	assert.Equal(t, []uint8{169, 169, 169}, rgbVal)

	// #3b82f6 = rgb(59, 130, 246) = hsl(217, 91%, 60%)
	rgbVal = HslIntToRgb(217, 91, 60)
	assert.Eq(t, []uint8{60, 131, 246}, rgbVal) // 计算后始终会有误差

	// RgbToHslInt
	hslVal := RgbToHslInt(rgbVal[0], rgbVal[1], rgbVal[2])
	assert.Equal(t, []int{217, 91, 60}, hslVal)
}

func TestRgbToHsl(t *testing.T) {
	hslFVal := RgbToHsl(57, 187, 226)
	fmt.Println("rgb: 57, 187, 226 => hsl:", hslFVal)
	assert.NotEmpty(t, hslFVal)

	hslFVal = RgbToHsl(255, 100, 50)
	fmt.Println("rgb: 255, 100, 50 => hsl:", hslFVal)
	assert.NotEmpty(t, hslFVal)

	hslFVal = RgbToHsl(50, 255, 100)
	fmt.Println("rgb: 50, 255, 100 => hsl:", hslFVal)
	assert.NotEmpty(t, hslFVal)

	hslFVal = RgbToHsl(100, 50, 255)
	fmt.Println("rgb: 100, 50, 255 => hsl:", hslFVal)
	assert.NotEmpty(t, hslFVal)

	// 红色 RGB(255,0,0) -> HSL(0.00, 1.00, 0.50)
	// 绿色 RGB(0,255,0) -> HSL(120.00, 1.00, 0.50)
	// 蓝色 RGB(0,0,255) -> HSL(240.00, 1.00, 0.50)
	// 灰色 RGB(128,128,128) -> HSL(0.00, 0.00, 0.50)
	// 黄色 RGB(255,255,0) -> HSL(60.00, 1.00, 0.50)
	tests := []struct {
		rgb []uint8
		hsl []int
	}{
		{[]uint8{255, 0, 0}, []int{0, 100, 50}},
		{[]uint8{0, 255, 0}, []int{120, 100, 50}},
		{[]uint8{0, 0, 255}, []int{240, 100, 50}},
		{[]uint8{128, 128, 128}, []int{0, 0, 50}},
	}
	for _, test := range tests {
		hslInts := RgbToHslInt(test.rgb[0], test.rgb[1], test.rgb[2])
		assert.Equal(t, test.hsl, hslInts, "input rgb: %v", test.rgb)
	}

	// #3b82f6 = rgb(59, 130, 246) = hsl(217, 91%, 60%)
	intVals := RgbStrToHslInts("rgb(59, 130, 246)")
	assert.NotNil(t, intVals)
	// fmt.Println("rgb(59, 130, 246) => hsl:", intVals)
	assert.Equal(t, []int{217, 91, 60}, intVals)

	intVals = RgbStrToHslInts("59, 130, 246")
	assert.NotNil(t, intVals)
	assert.Equal(t, []int{217, 91, 60}, intVals)
	// fmt.Println("rgb(59, 130, 246) => hsl:", intVals)

	// error cases
	intVals = RgbStrToHslInts("59, 246")
	assert.Nil(t, intVals)
	intVals = RgbStrToHslInts("abc, 130, 246")
	assert.Nil(t, intVals)
	intVals = RgbStrToHslInts("35, def, 246")
	assert.Nil(t, intVals)
	intVals = RgbStrToHslInts("59, 130, fda")
	assert.Nil(t, intVals)
}

func TestRgbHsvConv(t *testing.T) {
	// 示例1：纯红色 (255, 0, 0) -> HSV(0.00, 1.00, 1.00)
	h, s, v := RGBToHSV(255, 0, 0)
	assert.Eq(t, "(0.00, 1.00, 1.00)", fmt.Sprintf("(%.2f, %.2f, %.2f)", h, s, v))
	r, g, b := HSVToRGB(h, s, v)
	assert.Eq(t, "(255, 0, 0)", fmt.Sprintf("(%d, %d, %d)", r, g, b)) // TODO use assert.StrEq
	hsvInts := RGBToHSVInts(255, 0, 0)
	assert.Eq(t, "[0 100 100]", fmt.Sprint(hsvInts))
	hsvF64s := RgbToHsvSlice(255, 0, 0)
	assert.Eq(t, "[0 1 1]", fmt.Sprint(hsvF64s))

	// 示例2：纯绿色 (0, 255, 0) -> HSV(120.00, 1.00, 1.00)
	h, s, v = RGBToHSV(0, 255, 0)
	assert.Eq(t, "(120.00, 1.00, 1.00)", fmt.Sprintf("(%.2f, %.2f, %.2f)", h, s, v))
	r, g, b = HSVToRGB(h, s, v)
	assert.Eq(t, "(0, 255, 0)", fmt.Sprintf("(%d, %d, %d)", r, g, b))
	rgbInts := HSVIntToRGBInts(120, 100, 100)
	assert.Eq(t, "[0 255 0]", fmt.Sprint(rgbInts))

	// 示例3：纯蓝色 (0, 0, 255) -> HSV(240.00, 1.00, 1.00)
	h, s, v = RGBToHSV(0, 0, 255)
	assert.Eq(t, "(240.00, 1.00, 1.00)", fmt.Sprintf("(%.2f, %.2f, %.2f)", h, s, v))
	r, g, b = HSVToRGB(h, s, v)
	assert.Eq(t, "(0, 0, 255)", fmt.Sprintf("(%d, %d, %d)", r, g, b))

	// 示例4：黄色 (255, 255, 0) -> HSV(60.00, 1.00, 1.00)
	h, s, v = RGBToHSV(255, 255, 0)
	assert.Eq(t, "(60.00, 1.00, 1.00)", fmt.Sprintf("(%.2f, %.2f, %.2f)", h, s, v))
	r, g, b = HSVToRGB(h, s, v)
	assert.Eq(t, "(255, 255, 0)", fmt.Sprintf("(%d, %d, %d)", r, g, b))

	// 示例5：紫色 (128, 0, 128) -> HSV(300.00, 1.00, 0.50)
	h, s, v = RGBToHSV(128, 0, 128)
	assert.Eq(t, "(300.00, 1.00, 0.50)", fmt.Sprintf("(%.2f, %.2f, %.2f)", h, s, v))
	r, g, b = HSVToRGB(h, s, v)
	assert.Eq(t, "(128, 0, 128)", fmt.Sprintf("(%d, %d, %d)", r, g, b))

	// 示例6：黑色 (0, 0, 0) -> HSV(0.00, 0.00, 0.00)
	h, s, v = RGBToHSV(0, 0, 0)
	assert.Eq(t, "(0.00, 0.00, 0.00)", fmt.Sprintf("(%.2f, %.2f, %.2f)", h, s, v))
	r, g, b = HSVToRGB(h, s, v)
	assert.Eq(t, "(0, 0, 0)", fmt.Sprintf("(%d, %d, %d)", r, g, b))

	// 示例7：白色 (255, 255, 255) -> HSV(0.00, 0.00, 1.00)
	h, s, v = RGBToHSV(255, 255, 255)
	assert.Eq(t, "(0.00, 0.00, 1.00)", fmt.Sprintf("(%.2f, %.2f, %.2f)", h, s, v))
	r, g, b = HSVToRGB(h, s, v)
	assert.Eq(t, "(255, 255, 255)", fmt.Sprintf("(%d, %d, %d)", r, g, b))
}