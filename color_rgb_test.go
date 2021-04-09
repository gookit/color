package color

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*************************************************************
 * test rgb color
 *************************************************************/

func TestRGBColor(t *testing.T) {
	buf := forceOpenColorRender()
	defer resetColorRender()
	is := assert.New(t)

	// empty
	c := RGBColor{3: 99}
	is.True(c.IsEmpty())
	is.Equal("", c.String())

	// bg
	c = RGB(204, 204, 204, true)
	is.False(c.IsEmpty())
	is.Equal("48;2;204;204;204", c.FullCode())
	is.Equal("48;2;204;204;204", c.String())

	// fg
	c = RGB(204, 204, 204)
	is.False(c.IsEmpty())
	is.Equal("38;2;204;204;204", c.FullCode())
	is.Equal("38;2;204;204;204", c.String())

	// RGBColor.Sprint
	str := c.Sprint("msg")
	is.Equal("\x1b[38;2;204;204;204mmsg\x1b[0m", str)

	// RGBColor.Sprintf
	str = c.Sprintf("msg")
	is.Equal("\x1b[38;2;204;204;204mmsg\x1b[0m", str)
	is.Equal("cccccc", c.Hex())
	is.Equal("204;204;204", c.Code())
	is.Equal("38;2;204;204;204", c.FullCode())
	is.Equal("[204 204 204]", fmt.Sprint(c.Values()))

	// RGBColor.Print
	c.Print("msg")
	str = buf.String()
	buf.Reset()
	is.Equal("\x1b[38;2;204;204;204mmsg\x1b[0m", str)

	// RGBColor.Printf
	c.Printf("m%s", "sg")
	str = buf.String()
	buf.Reset()
	is.Equal("\x1b[38;2;204;204;204mmsg\x1b[0m", str)

	// RGBColor.Println
	c.Println("msg")
	str = buf.String()
	buf.Reset()
	is.Equal("\x1b[38;2;204;204;204mmsg\x1b[0m\n", str)
}

func TestRGBFromString(t *testing.T) {
	forceOpenColorRender()
	defer resetColorRender()
	is := assert.New(t)

	c := RGBFromString("170,187,204")
	is.Equal("\x1b[38;2;170;187;204mmsg\x1b[0m", c.Sprint("msg"))

	c = RGBFromString("170,187,204", true)
	is.Equal("\x1b[48;2;170;187;204mmsg\x1b[0m", c.Sprint("msg"))

	c = RGBFromString("170,187,")
	is.Equal("msg", c.Sprint("msg"))

	c = RGBFromString("")
	is.Equal("msg", c.Sprint("msg"))

	c = RGBFromString("170,187,error")
	is.Equal("msg", c.Sprint("msg"))
}

func TestHexToRGB(t *testing.T) {
	is := assert.New(t)
	c := HEX("ccc") // rgb: [204 204 204]
	is.False(c.IsEmpty())
	is.Equal("38;2;204;204;204", c.String())
	is.Equal("cccccc", c.Hex())

	c = HEX("aabbcc") // rgb: [170 187 204]
	is.Equal("38;2;170;187;204", c.String())
	is.Equal("aabbcc", c.Hex())

	c = Hex("#aabbcc") // rgb: [170 187 204]
	is.Equal("38;2;170;187;204", c.String())

	c = HEX("0xad99c0") // rgb: [170 187 204]
	is.Equal("38;2;173;153;192", c.String())

	c = HEX(" ")
	is.True(c.IsEmpty())
	is.Equal("", c.String())

	c = HEX("!#$bbcc")
	is.Equal("", c.String())

	c = HEX("#invalid")
	is.Equal("", c.String())

	c = HEX("invalid code")
	is.Equal("", c.String())
}

func TestRGBStyle(t *testing.T) {
	is := assert.New(t)

	ForceColor()
	fg := RGB(20, 144, 234)
	bg := RGB(234, 78, 23)

	s := &RGBStyle{}
	is.True(s.IsEmpty())
	is.Equal("", s.String())

	s.Set(fg, bg)
	is.False(s.IsEmpty())
	is.Equal("38;2;20;144;234;48;2;234;78;23", s.String())

	s = &RGBStyle{}
	s.Set(fg, bg, OpUnderscore)
	is.False(s.IsEmpty())
	is.Equal("38;2;20;144;234;48;2;234;78;23;4", s.FullCode())
	is.Equal("38;2;20;144;234;48;2;234;78;23;4", s.String())

	s.SetOpts(Opts{OpBold, OpBlink})
	is.Equal("38;2;20;144;234;48;2;234;78;23;1;5", s.Code())
	is.Equal("38;2;20;144;234;48;2;234;78;23;1;5", s.String())

	s.AddOpts(OpItalic)
	is.Equal("38;2;20;144;234;48;2;234;78;23;1;5;3", s.String())

	// NewRGBStyle
	s = NewRGBStyle(RGB(20, 144, 234))
	is.False(s.IsEmpty())
	is.Equal("38;2;20;144;234", s.String())

	s = NewRGBStyle(RGB(20, 144, 234), RGB(234, 78, 23))
	is.False(s.IsEmpty())
	is.Equal("38;2;20;144;234;48;2;234;78;23", s.String())

	// HEXStyle
	s = HEXStyle("555", "eee")
	is.False(s.IsEmpty())
	is.Equal("38;2;85;85;85;48;2;238;238;238", s.String())

	// RGBStyleFromString
	s = RGBStyleFromString("20, 144, 234", "234, 78, 23")
	is.False(s.IsEmpty())
	is.Equal("38;2;20;144;234;48;2;234;78;23", s.String())

	// RGBColor.Sprint
	is.Equal("\x1b[38;2;20;144;234;48;2;234;78;23mmsg\x1b[0m", s.Sprint("msg"))
	// RGBColor.Sprintf
	is.Equal("\x1b[38;2;20;144;234;48;2;234;78;23mmsg\x1b[0m", s.Sprintf("m%s", "sg"))

	s.Println("hello, this is use RGB color")
	fmt.Println("\x1b[38;2;20;144;234;48;2;234;78;23mTEXT\x1b[0m")
	// add option: OpItalic
	fmt.Println("\x1b[38;2;20;144;234;48;2;234;78;23;3mTEXT\x1b[0m")

	buf := forceOpenColorRender()
	defer resetColorRender()

	// RGBColor.Print
	s.Print("msg")
	str := buf.String()
	buf.Reset()
	is.Equal("\x1b[38;2;20;144;234;48;2;234;78;23mmsg\x1b[0m", str)

	// RGBColor.Printf
	s.Printf("m%s", "sg")
	str = buf.String()
	buf.Reset()
	is.Equal("\x1b[38;2;20;144;234;48;2;234;78;23mmsg\x1b[0m", str)

	// RGBColor.Println
	s.Println("msg")
	str = buf.String()
	buf.Reset()
	is.Equal("\x1b[38;2;20;144;234;48;2;234;78;23mmsg\x1b[0m\n", str)
}

func TestPrintRGBColor(t *testing.T) {
	RGB(30, 144, 255).Println("message. use RGB number")
	HEX("#1976D2").Println("blue-darken")
	RGBStyleFromString("213,0,0").Println("red-accent. use RGB number")
	// foreground: eee, background: D50000
	HEXStyle("eee", "D50000").Println("deep-purple color")
}

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
	assert.Equal(t, "36", rgb.Color().String())

	rgb = RGB(57, 187, 226, true)
	assert.Equal(t, "46", rgb.C16().String())
	assert.Equal(t, "46", rgb.Color().String())
}
