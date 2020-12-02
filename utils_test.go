package color

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexToRgb(t *testing.T) {
	tests := []struct {
		given string
		want  []int
	}{
		{"666", []int{102, 102, 102}},
		{"ccc", []int{204, 204, 204}},
		{"#abc", []int{170, 187, 204}},
		{"#aa99cd", []int{170, 153, 205}},
	}

	for _, item := range tests {
		assert.Equal(t, HexToRgb(item.given), item.want)
		assert.Equal(t, HexToRGB(item.given), item.want)
		assert.Equal(t, Hex2rgb(item.given), item.want)
	}

	assert.Len(t, HexToRgb(""), 0)
	assert.Len(t, HexToRgb("13"), 0)
}

func TestRgbToHex(t *testing.T) {
	tests := []struct {
		want  string
		given []int
	}{
		{"666666", []int{102, 102, 102}},
		{"cccccc", []int{204, 204, 204}},
		{"aabbcc", []int{170, 187, 204}},
		{"aa99cd", []int{170, 153, 205}},
	}

	for _, item := range tests {
		assert.Equal(t, RgbToHex(item.given), item.want)
		assert.Equal(t, Rgb2hex(item.given), item.want)
	}
}

func TestRgbToAnsi(t *testing.T) {
	tests := []struct {
		want uint8
		rgb  []uint8
		isBg bool
	}{
		{40, []uint8{102, 102, 102}, true},
		{37, []uint8{204, 204, 204}, false},
		{47, []uint8{170, 78, 204}, true},
		{37, []uint8{170, 153, 245}, false},
		{30, []uint8{127, 127, 127}, false},
		{40, []uint8{127, 127, 127}, true},
		{90, []uint8{128, 128, 128}, false},
		{97, []uint8{34, 56, 255}, false},
		{31, []uint8{134, 56, 56}, false},
		{30, []uint8{0, 0, 0}, false},
		{40, []uint8{0, 0, 0}, true},
		{97, []uint8{255, 255, 255}, false},
		{107, []uint8{255, 255, 255}, true},
	}

	for _, item := range tests {
		r, g, b := item.rgb[0], item.rgb[1], item.rgb[2]

		assert.Equal(
			t,
			item.want,
			RgbToAnsi(r, g, b, item.isBg),
			fmt.Sprint("rgb=", item.rgb, ", is bg? ", item.isBg),
		)
		assert.Equal(t, item.want, Rgb2ansi(r, g, b, item.isBg))
	}
}
