package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testRgbToC256Color(t *testing.T, name string, c RGBColor, expected uint8) {
	t.Log("RGB Color:", c.Sprint(name))
	t.Log("256 Color:", c.C256().Sprint(name))
	actual := c.C256().Value()
	if actual != expected {
		t.Errorf("%s not converted correctly: expected %v, actual %v", name, actual, expected)
	}
}

func TestRGBStyle_SetOpts(t *testing.T) {
	s := NewRGBStyle(RGB(234, 78, 23), RGB(20, 144, 234))
	s.Println("rgb style message")

	s.SetOpts(Opts{OpItalic, OpBold, OpUnderscore})
	s.Println("RGB style message with options")
}

func TestRgbToC256(t *testing.T) {
	testRgbToC256Color(t, "white", RGB(255, 255, 255), 15)
	testRgbToC256Color(t, "red", RGB(255, 0, 0), 9)
	testRgbToC256Color(t, "yellow", RGB(255, 255, 0), 11)
	testRgbToC256Color(t, "greenBg", RGB(0, 255, 0, true), 10)
	testRgbToC256Color(t, "blueBg", RGB(0, 0, 255, true), 12)
	testRgbToC256Color(t, "light blue", RGB(57, 187, 226), 74)
}

func TestRgbToC256Background(t *testing.T) {
	white := Rgb(255, 255, 255)
	whiteBg := Bit24(255, 255, 255, true)
	whiteFg := RGB(255, 255, 255, false)
	if white.C256().String() != whiteFg.C256().String() {
		t.Error("standard color didn't match foreground color")
	}
	if white.C256().String() == whiteBg.C256().String() {
		t.Error("standard color matched background color")
	}
	prefix := whiteBg.C256().String()[:3]
	if prefix != "48;" {
		t.Errorf("background color didn't have background prefix: %v", prefix)
	}
}

func TestRGBColor_C16(t *testing.T) {
	rgb := RGB(57, 187, 226)
	assert.Equal(t, "36", rgb.C16().String())

	rgb = RGB(57, 187, 226, true)
	assert.Equal(t, "46", rgb.C16().String())
}
