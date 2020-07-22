package color

import (
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
