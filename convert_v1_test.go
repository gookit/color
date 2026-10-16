package color

import (
	"testing"

	"github.com/gookit/assert"
)

func TestC256ToRgbV1ColorCube(t *testing.T) {
	levels := []uint8{0, 42, 85, 127, 170, 212}
	for r, red := range levels {
		for g, green := range levels {
			for b, blue := range levels {
				index := uint8(16 + 36*r + 6*g + b)
				assert.Equal(t, []uint8{red, green, blue}, C256ToRgbV1(index), "palette index %d", index)
			}
		}
	}
}

func TestC256ToRgbV1BaseColors(t *testing.T) {
	colors := [][]uint8{
		{0, 0, 0}, {170, 0, 0}, {0, 170, 0}, {170, 170, 0},
		{0, 0, 170}, {170, 0, 170}, {0, 170, 170}, {170, 170, 170},
		{85, 85, 85}, {255, 85, 85}, {85, 255, 85}, {255, 255, 85},
		{85, 85, 255}, {255, 85, 255}, {85, 255, 255}, {255, 255, 255},
	}
	for index, want := range colors {
		assert.Equal(t, want, C256ToRgbV1(uint8(index)), "palette index %d", index)
	}
}

func TestC256ToRgbV1Grayscale(t *testing.T) {
	levels := []uint8{
		8, 18, 28, 38, 48, 58, 68, 78, 88, 98, 108, 118,
		128, 138, 148, 158, 168, 178, 188, 198, 208, 218, 228, 238,
	}
	for index, level := range levels {
		assert.Equal(t, []uint8{level, level, level}, C256ToRgbV1(uint8(232+index)), "gray level %d", index)
	}
}
