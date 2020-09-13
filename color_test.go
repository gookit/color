package color

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Example() {
	// quick use like fmt.Print*
	Red.Println("Simple to use color")
	Green.Print("Simple to use color")
	Cyan.Printf("Simple to use %s\n", "color")
	Gray.Printf("Simple to use %s\n", "color")
	Blue.Printf("Simple to use %s\n", "color")
	Yellow.Printf("Simple to use %s\n", "color")
	Magenta.Printf("Simple to use %s\n", "color")

	// use like func
	red := FgRed.Render
	green := FgGreen.Render
	fmt.Printf("%s line %s library\n", red("Command"), green("color"))

	// custom color
	New(FgWhite, BgBlack).Println("custom color style")

	// can also:
	Style{FgCyan, OpBold}.Println("custom color style")

	// internal theme/style:
	Info.Tips("message")
	Info.Prompt("message")
	Info.Println("info message")
	Warn.Println("warning message")
	Error.Println("error message")
	Danger.Println("danger message")

	// use style tag
	Print("<suc>he</><comment>llo</>, <cyan>wel</><red>come</>\n")

	// apply a style tag
	Tag("info").Println("info style text")

	// prompt message
	Info.Prompt("prompt style message")
	Warn.Prompt("prompt style message")

	// tips message
	Info.Tips("tips style message")
	Warn.Tips("tips style message")
}

/*************************************************************
 * test global methods
 *************************************************************/

func TestSet(t *testing.T) {
	is := assert.New(t)

	// set
	num, err := Set(FgGreen)
	fmt.Println("test color.Set() on OS:", runtime.GOOS)
	is.True(num > 0)
	is.NoError(err)
	_, err = Reset()
	is.NoError(err)

	// disable
	old := Disable()
	num, err = Set(FgGreen)
	is.Nil(err)
	is.Equal(0, num)

	num, err = Reset()
	is.Nil(err)
	is.Equal(0, num)
	Enable = old // revert

	if runtime.GOOS == "windows" {
		fd := uintptr(syscall.Stdout)
		// if run test by goland, will return false
		if IsTerminal(int(fd)) {
			fmt.Println("- IsTerminal return TRUE")
		} else {
			fmt.Println("- IsTerminal return FALSE")
		}
	} else {
		is.False(IsLikeInCmd())
		is.Empty(GetErrors())
	}

	// set
	rewriteStdout()
	num, err = Set(FgGreen)
	str := restoreStdout()

	is.True(num > 0)
	is.NoError(err)
	is.Equal("\x1b[32m", str)

	// unset
	rewriteStdout()
	_, err = Reset()
	str = restoreStdout()
	is.NoError(err)
	is.Equal("\x1b[0m", str)
}

func TestRenderCode(t *testing.T) {
	// force open color render for testing
	oldVal = ForceColor()
	defer resetColorRender()

	is := assert.New(t)

	str := RenderCode("36;1", "Hi,", "babe")
	is.Equal("\x1b[36;1mHi,babe\x1b[0m", str)

	str = RenderWithSpaces("", "Hi,", "babe")
	is.Equal("Hi, babe", str)

	str = RenderWithSpaces("36;1", "Hi,", "babe")
	is.Equal("\x1b[36;1mHi, babe\x1b[0m", str)

	str = RenderCode("36;1", "Ab")
	is.Equal("\x1b[36;1mAb\x1b[0m", str)

	str = RenderCode("36;1")
	is.Equal("", str)

	Disable()
	str = RenderCode("36;1", "Te", "xt")
	is.Equal("Text", str)

	str = RenderWithSpaces("36;1", "Te", "xt")
	is.Equal("Te xt", str)
	Enable = true

	// RenderString
	str = RenderString("36;1", "Text")
	is.Equal("\x1b[36;1mText\x1b[0m", str)
	str = RenderString("", "Text")
	is.Equal("Text", str)
	str = RenderString("36;1", "")
	is.Equal("", str)

	Disable()
	str = RenderString("36;1", "Text")
	is.Equal("Text", str)
	Enable = true

	Disable()
	str = RenderString("36;1", "Text")
	is.Equal("Text", str)
	Enable = true
}

func TestClearCode(t *testing.T) {
	art := assert.New(t)
	art.Equal("Text", ClearCode("\033[36;1mText\x1b[0m"))
	// 8bit
	art.Equal("Text", ClearCode("\x1b[38;5;242mText\x1b[0m"))
	// 24bit
	art.Equal("Text", ClearCode("\x1b[38;2;30;144;255mText\x1b[0m"))
	art.Equal("Text other", ClearCode("\033[36;1mText\x1b[0m other"))
}

/*************************************************************
 * test printer
 *************************************************************/

func TestPrinter(t *testing.T) {
	buf := forceOpenColorRender()
	defer resetColorRender()
	at := assert.New(t)

	p := NewPrinter("48;5;132")

	// Color256.Sprint
	str := p.Sprint("msg")
	at.Equal("\x1b[48;5;132mmsg\x1b[0m", str)
	// Color256.Sprintf
	str = p.Sprintf("msg")
	at.Equal("\x1b[48;5;132mmsg\x1b[0m", str)

	at.False(p.IsEmpty())
	at.Equal("48;5;132", p.String())

	// Color256.Print
	p.Print("MSG")
	str = buf.String()
	buf.Reset()
	at.Equal("\x1b[48;5;132mMSG\x1b[0m", str)

	// Color256.Printf
	p.Printf("A %s", "MSG")
	str = buf.String()
	buf.Reset()
	at.Equal("\x1b[48;5;132mA MSG\x1b[0m", str)

	// Color256.Println
	p.Println("MSG")
	str = buf.String()
	buf.Reset()
	at.Equal("\x1b[48;5;132mMSG\x1b[0m\n", str)
}

/*************************************************************
 * test 16 color
 *************************************************************/

func TestColor16(t *testing.T) {
	buf := forceOpenColorRender()
	defer resetColorRender()
	at := assert.New(t)

	at.True(Bold.IsValid())
	r := Bold.Render("text")
	at.Equal("\x1b[1mtext\x1b[0m", r)
	r = LightYellow.Text("text")
	at.Equal("\x1b[93mtext\x1b[0m", r)
	r = LightWhite.Sprint("text")
	at.Equal("\x1b[97mtext\x1b[0m", r)
	r = White.Render("test", "spaces")
	at.Equal("\x1b[37mtestspaces\x1b[0m", r)
	r = Black.Renderln("test", "spaces")
	at.Equal("\x1b[30mtest spaces\x1b[0m", r)

	str := Red.Sprintf("A %s", "MSG")
	at.Equal("\x1b[31mA MSG\x1b[0m", str)

	// Color.Print
	FgGray.Print("MSG")
	str = buf.String()
	at.Equal("\x1b[90mMSG\x1b[0m", str)
	buf.Reset()

	// Color.Printf
	BgGray.Printf("A %s", "MSG")
	str = buf.String()
	at.Equal("\x1b[100mA MSG\x1b[0m", str)
	buf.Reset()

	// Color.Println
	LightMagenta.Println("MSG")
	str = buf.String()
	at.Equal("\x1b[95mMSG\x1b[0m\n", str)
	buf.Reset()

	LightMagenta.Println()
	str = buf.String()
	at.Equal("\n", str)
	buf.Reset()

	LightCyan.Print("msg")
	LightRed.Printf("m%s", "sg")
	LightGreen.Println("msg")
	str = buf.String()
	at.Equal("\x1b[96mmsg\x1b[0m\x1b[91mmsg\x1b[0m\x1b[92mmsg\x1b[0m\n", str)
	buf.Reset()

	// Color.Darken
	blue := LightBlue.Darken()
	at.Equal(94, int(LightBlue))
	at.Equal(34, int(blue))
	c := Color(120).Darken()
	at.Equal(120, int(c))

	// Color.Light
	lightCyan := Cyan.Light()
	at.Equal(36, int(Cyan))
	at.Equal(96, int(lightCyan))
	c = Color(120).Light()
	at.Equal(120, int(c))

	// Colors vars
	_, ok := FgColors["red"]
	at.True(ok)
	_, ok = ExFgColors["lightRed"]
	at.True(ok)
	_, ok = BgColors["red"]
	at.True(ok)
	_, ok = ExBgColors["lightRed"]
	at.True(ok)
}

func TestPrintBasicColor(t *testing.T) {
	fmt.Println("Foreground colors:")
	for name, c := range FgColors {
		c.Print(" ", name, " ")
	}

	fmt.Println("\nBackground colors:")
	for name, c := range BgColors {
		c.Print(" ", name, " ")
	}

	fmt.Println("\nBasic Options:")
	for name, c := range Options {
		c.Print(" ", name, " ")
	}

	fmt.Println("\nExtra foreground colors:")
	for name, c := range ExFgColors {
		c.Print(" ", name, " ")
	}

	fmt.Println("\nExtra background colors:")
	for name, c := range ExBgColors {
		c.Print(" ", name, " ")
	}

	fmt.Println()
	fmt.Println()
}

/*************************************************************
 * test 256 color
 *************************************************************/

func TestColor256(t *testing.T) {
	buf := forceOpenColorRender()
	defer resetColorRender()

	is := assert.New(t)

	// empty
	c := Color256{1: 99}
	is.True(c.IsEmpty())
	is.Equal("", c.String())

	// fg
	c = Bit8(132)
	is.False(c.IsEmpty())
	is.Equal(uint8(132), c.Value())
	is.Equal("38;5;132", c.String())

	// Color256.Sprint
	str := c.Sprint("msg")
	is.Equal("\x1b[38;5;132mmsg\x1b[0m", str)
	// Color256.Sprintf
	str = c.Sprintf("msg")
	is.Equal("\x1b[38;5;132mmsg\x1b[0m", str)

	// bg
	c = Bit8(132, true)
	is.False(c.IsEmpty())
	is.Equal("48;5;132", c.String())

	c = C256(132)
	// Color256.Print
	c.Print("MSG")
	str = buf.String()
	buf.Reset()
	is.Equal("\x1b[38;5;132mMSG\x1b[0m", str)

	// Color256.Printf
	c.Printf("A %s", "MSG")
	str = buf.String()
	buf.Reset()
	is.Equal("\x1b[38;5;132mA MSG\x1b[0m", str)

	// Color256.Println
	c.Println("MSG", "TEXT")
	str = buf.String()
	buf.Reset()
	is.Equal("\x1b[38;5;132mMSG TEXT\x1b[0m\n", str)
}

func TestStyle256(t *testing.T) {
	buf := forceOpenColorRender()
	defer resetColorRender()

	is := assert.New(t)
	// empty
	s := S256()
	is.Equal("", s.String())
	is.Equal("MSG", s.Sprint("MSG"))

	// only fg
	s = S256(132)
	is.Equal("38;5;132", s.String())
	is.Equal("\x1b[38;5;132mMSG\x1b[0m", s.Sprint("MSG"))
	is.Equal("\x1b[38;5;132mMSG\x1b[0m", s.Sprintf("%s", "MSG"))

	// only bg
	s = S256(132)
	is.Equal("38;5;132", s.String())
	is.Equal("\x1b[38;5;132mMSG\x1b[0m", s.Sprint("MSG"))

	// fg and bg
	s = S256(132, 23)
	is.Equal("38;5;132;48;5;23", s.String())
	is.Equal("\x1b[38;5;132;48;5;23mMSG\x1b[0m", s.Sprint("MSG"))
	s = S256().Set(132, 23)
	is.Equal("38;5;132;48;5;23", s.String())
	is.Equal("\x1b[38;5;132;48;5;23mMSG\x1b[0m", s.Sprint("MSG"))
	s = S256().SetFg(132).SetBg(23)
	is.Equal("38;5;132;48;5;23", s.String())
	is.Equal("\x1b[38;5;132;48;5;23mMSG\x1b[0m", s.Sprint("MSG"))

	s = S256(132)

	// Color256.Print
	s.Print("MSG")
	str := buf.String()
	buf.Reset()
	is.Equal("\x1b[38;5;132mMSG\x1b[0m", str)

	// Color256.Printf
	s.Printf("A %s", "MSG")
	str = buf.String()
	buf.Reset()
	is.Equal("\x1b[38;5;132mA MSG\x1b[0m", str)

	// Color256.Println
	s.Println("MSG")
	str = buf.String()
	buf.Reset()
	is.Equal("\x1b[38;5;132mMSG\x1b[0m\n", str)
}

func TestPrint256color(t *testing.T) {
	fmt.Printf("\n%-50s24th Order Grayscale Color\n", " ")

	var fg uint8 = 255
	for i := range []int{23: 0} { // // 232-255：从黑到白的24阶灰度色
		if i < 12 {
			fg = 255
		} else {
			fg = 0
		}

		i += 232
		S256(fg, uint8(i)).Printf(" %-4d", i)
	}
	fmt.Println()
	fmt.Println()
}

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
	is.Equal("48;2;204;204;204", c.String())

	// fg
	c = RGB(204, 204, 204)
	is.False(c.IsEmpty())
	is.Equal("38;2;204;204;204", c.String())

	// RGBColor.Sprint
	str := c.Sprint("msg")
	is.Equal("\x1b[38;2;204;204;204mmsg\x1b[0m", str)

	// RGBColor.Sprintf
	str = c.Sprintf("msg")
	is.Equal("\x1b[38;2;204;204;204mmsg\x1b[0m", str)
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

	c = HEX("aabbcc") // rgb: [170 187 204]
	is.Equal("38;2;170;187;204", c.String())

	c = HEX("#aabbcc") // rgb: [170 187 204]
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
	buf := forceOpenColorRender()
	defer resetColorRender()
	at := assert.New(t)

	s := &RGBStyle{}
	at.True(s.IsEmpty())
	at.Equal("", s.String())
	s.Set(RGB(20, 144, 234), RGB(234, 78, 23))
	at.False(s.IsEmpty())
	at.Equal("38;2;20;144;234;48;2;234;78;23", s.String())

	// NewRGBStyle
	s = NewRGBStyle(RGB(20, 144, 234))
	at.False(s.IsEmpty())
	at.Equal("38;2;20;144;234", s.String())

	s = NewRGBStyle(RGB(20, 144, 234), RGB(234, 78, 23))
	at.False(s.IsEmpty())
	at.Equal("38;2;20;144;234;48;2;234;78;23", s.String())

	// HEXStyle
	s = HEXStyle("555", "eee")
	at.False(s.IsEmpty())
	at.Equal("38;2;85;85;85;48;2;238;238;238", s.String())

	// RGBStyleFromString
	s = RGBStyleFromString("20, 144, 234", "234, 78, 23")
	at.False(s.IsEmpty())
	at.Equal("38;2;20;144;234;48;2;234;78;23", s.String())

	// RGBColor.Sprint
	at.Equal("\x1b[38;2;20;144;234;48;2;234;78;23mmsg\x1b[0m", s.Sprint("msg"))
	// RGBColor.Sprintf
	at.Equal("\x1b[38;2;20;144;234;48;2;234;78;23mmsg\x1b[0m", s.Sprintf("m%s", "sg"))

	// RGBColor.Print
	s.Print("msg")
	str := buf.String()
	buf.Reset()
	at.Equal("\x1b[38;2;20;144;234;48;2;234;78;23mmsg\x1b[0m", str)

	// RGBColor.Printf
	s.Printf("m%s", "sg")
	str = buf.String()
	buf.Reset()
	at.Equal("\x1b[38;2;20;144;234;48;2;234;78;23mmsg\x1b[0m", str)

	// RGBColor.Println
	s.Println("msg")
	str = buf.String()
	buf.Reset()
	at.Equal("\x1b[38;2;20;144;234;48;2;234;78;23mmsg\x1b[0m\n", str)
}

func TestPrintRGBColor(t *testing.T) {
	RGB(30, 144, 255).Println("message. use RGB number")
	HEX("#1976D2").Println("blue-darken")
	RGBStyleFromString("213,0,0").Println("red-accent. use RGB number")
	// foreground: eee, background: D50000
	HEXStyle("eee", "D50000").Println("deep-purple color")
}

func TestUtilFuncs(t *testing.T) {
	is := assert.New(t)

	// IsConsole
	is.True(IsConsole(os.Stdin))
	is.True(IsConsole(os.Stdout))
	is.True(IsConsole(os.Stderr))
	is.False(IsConsole(&bytes.Buffer{}))
	ff, err := os.OpenFile(".travis.yml", os.O_WRONLY, 0)
	is.NoError(err)
	is.False(IsConsole(ff))

	// IsMSys
	oldVal := os.Getenv("MSYSTEM")
	is.NoError(os.Setenv("MSYSTEM", "MINGW64"))
	is.True(IsMSys())
	is.NoError(os.Unsetenv("MSYSTEM"))
	is.False(IsMSys())
	_ = os.Setenv("MSYSTEM", oldVal)

	// IsSupport256Color
	oldVal = os.Getenv("TERM")
	_ = os.Unsetenv("TERM")
	is.False(IsSupportColor())
	is.False(IsSupport256Color())

	// ConEmuANSI
	mockEnvValue("ConEmuANSI", "ON", func(_ string) {
		is.True(IsSupportColor())
	})

	// ANSICON
	mockEnvValue("ANSICON", "189x2000 (189x43)", func(_ string) {
		is.True(IsSupportColor())
	})

	// "COLORTERM=truecolor"
	mockEnvValue("COLORTERM", "truecolor", func(_ string) {
		is.True(IsSupportTrueColor())
	})

	// TERM
	mockEnvValue("TERM", "screen-256color", func(_ string) {
		is.True(IsSupportColor())
	})

	// TERM
	mockEnvValue("TERM", "tmux-256color", func(_ string) {
		is.True(IsSupportColor())
	})

	// TERM
	mockEnvValue("TERM", "rxvt-unicode-256color", func(_ string) {
		is.True(IsSupportColor())
	})

	is.NoError(os.Setenv("TERM", "xterm-vt220"))
	is.True(IsSupportColor())
	// revert
	if oldVal != "" {
		is.NoError(os.Setenv("TERM", oldVal))
	} else {
		is.NoError(os.Unsetenv("TERM"))
	}
}

/*************************************************************
 * test helpers
 *************************************************************/

var oldVal bool

// force open color render for testing
func forceOpenColorRender() *bytes.Buffer {
	oldVal = isSupportColor
	isSupportColor = true

	// set output for test
	buf := new(bytes.Buffer)
	SetOutput(buf)

	return buf
}

func resetColorRender() {
	isSupportColor = oldVal
	// reset
	ResetOutput()
}

var oldStdout, newReader *os.File

// Usage:
// rewriteStdout()
// fmt.Println("Hello, playground")
// msg := restoreStdout()
func rewriteStdout() {
	oldStdout = os.Stdout
	r, w, _ := os.Pipe()
	newReader = r
	os.Stdout = w
}

func restoreStdout() string {
	if newReader == nil {
		return ""
	}

	// Notice: must close writer before read data
	// close now writer
	_ = os.Stdout.Close()
	// restore
	os.Stdout = oldStdout
	oldStdout = nil

	// read data
	out, _ := ioutil.ReadAll(newReader)

	// close reader
	_ = newReader.Close()
	newReader = nil

	return string(out)
}

// mockEnvValue will store old env value, set new val. will restore old value on end.
func mockEnvValue(key, val string, fn func(nv string)) {
	old := os.Getenv(key)
	err := os.Setenv(key, val)
	if err != nil {
		panic(err)
	}

	fn(os.Getenv(key))

	// if old is empty, unset key.
	if old == "" {
		err = os.Unsetenv(key)
	} else {
		err = os.Setenv(key, old)
	}
	if err != nil {
		panic(err)
	}
}
