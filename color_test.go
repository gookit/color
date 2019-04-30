package color

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Example() {
	// simple usage
	Cyan.Printf("Simple to use %s\n", "color")

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
	at := assert.New(t)

	// disable color
	old := Enable
	Disable()
	num, err := Set(FgGreen)
	at.Nil(err)
	at.Equal(0, num)

	num, err = Reset()
	at.Nil(err)
	at.Equal(0, num)
	Enable = old

	// set
	rewriteStdout()
	num, err = Set(FgGreen)
	str := restoreStdout()
	at.NoError(err)
	if isLikeInCmd {
		at.Equal("", str)
	} else {
		at.Equal("\x1b[32m", str)
	}
	_, _ = Reset()

	// unset
	rewriteStdout()
	_, err = Reset()
	str = restoreStdout()
	at.NoError(err)
	if isLikeInCmd {
		at.Equal("", str)
	} else {
		at.Equal("\x1b[0m", str)
	}

	if isLikeInCmd {
		// set
		rewriteStdout()
		_, err = Set(FgGreen)
		str = restoreStdout()
		at.NoError(err)
		at.Equal("", str)

		// unset
		rewriteStdout()
		_, err = Reset()
		str = restoreStdout()
		at.NoError(err)
		at.Equal("", str)
	}
}

func TestRenderCode(t *testing.T) {
	// force open color render for testing
	forceOpenColorRender()
	defer resetColorRender()

	at := assert.New(t)

	str := RenderCode("36;1", "Te", "xt")
	at.Equal("\x1b[36;1mText\x1b[0m", str)

	Disable()
	str = RenderCode("36;1", "Te", "xt")
	at.Equal("Text", str)
	Enable = true

	// RenderString
	str = RenderString("36;1", "Text")
	at.Equal("\x1b[36;1mText\x1b[0m", str)
	str = RenderString("", "Text")
	at.Equal("Text", str)
	str = RenderString("36;1", "")
	at.Equal("", str)

	Disable()
	str = RenderString("36;1", "Text")
	at.Equal("Text", str)
	Enable = true

	Disable()
	str = RenderString("36;1", "Text")
	at.Equal("Text", str)
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
	forceOpenColorRender()
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
	rewriteStdout()
	p.Print("MSG")
	str = restoreStdout()
	at.Equal("\x1b[48;5;132mMSG\x1b[0m", str)

	// Color256.Printf
	rewriteStdout()
	p.Printf("A %s", "MSG")
	str = restoreStdout()
	at.Equal("\x1b[48;5;132mA MSG\x1b[0m", str)

	// Color256.Println
	rewriteStdout()
	p.Println("MSG")
	str = restoreStdout()
	at.Equal("\x1b[48;5;132mMSG\x1b[0m\n", str)
}

/*************************************************************
 * test 16 color
 *************************************************************/

func TestColor16(t *testing.T) {
	forceOpenColorRender()
	defer resetColorRender()
	at := assert.New(t)

	at.True(Bold.IsValid())
	r := Bold.Render("text")
	at.Equal("\x1b[1mtext\x1b[0m", r)
	r = LightYellow.Text("text")
	at.Equal("\x1b[93mtext\x1b[0m", r)
	r = LightWhite.Sprint("text")
	at.Equal("\x1b[97mtext\x1b[0m", r)

	str := Red.Sprintf("A %s", "MSG")
	at.Equal("\x1b[31mA MSG\x1b[0m", str)

	// Color.Print
	rewriteStdout()
	FgGray.Print("MSG")
	str = restoreStdout()
	if isLikeInCmd {
		at.Equal("MSG", str)
	} else {
		at.Equal("\x1b[90mMSG\x1b[0m", str)
	}

	// Color.Printf
	rewriteStdout()
	BgGray.Printf("A %s", "MSG")
	str = restoreStdout()
	if isLikeInCmd {
		at.Equal("A MSG", str)
	} else {
		at.Equal("\x1b[100mA MSG\x1b[0m", str)
	}

	// Color.Println
	rewriteStdout()
	LightMagenta.Println("MSG")
	str = restoreStdout()
	if isLikeInCmd {
		at.Equal("MSG\n", str)
	} else {
		at.Equal("\x1b[95mMSG\x1b[0m\n", str)
	}

	if isLikeInCmd {
		rewriteStdout()
		LightCyan.Print("msg")
		LightRed.Printf("m%s", "sg")
		LightGreen.Println("msg")
		str = restoreStdout()
		at.Equal("msgmsgmsg\n", str)
	}

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

/*************************************************************
 * test 256 color
 *************************************************************/

func TestColor256(t *testing.T) {
	forceOpenColorRender()
	defer resetColorRender()

	at := assert.New(t)

	// empty
	c := Color256{1: 99}
	at.True(c.IsEmpty())
	at.Equal("", c.String())

	// fg
	c = Bit8(132)
	at.False(c.IsEmpty())
	at.Equal(uint8(132), c.Value())
	at.Equal("38;5;132", c.String())

	// Color256.Sprint
	str := c.Sprint("msg")
	at.Equal("\x1b[38;5;132mmsg\x1b[0m", str)
	// Color256.Sprintf
	str = c.Sprintf("msg")
	at.Equal("\x1b[38;5;132mmsg\x1b[0m", str)

	// bg
	c = Bit8(132, true)
	at.False(c.IsEmpty())
	at.Equal("48;5;132", c.String())

	c = C256(132)
	// Color256.Print
	rewriteStdout()
	c.Print("MSG")
	str = restoreStdout()
	at.Equal("\x1b[38;5;132mMSG\x1b[0m", str)

	// Color256.Printf
	rewriteStdout()
	c.Printf("A %s", "MSG")
	str = restoreStdout()
	at.Equal("\x1b[38;5;132mA MSG\x1b[0m", str)

	// Color256.Println
	rewriteStdout()
	c.Println("MSG")
	str = restoreStdout()
	at.Equal("\x1b[38;5;132mMSG\x1b[0m\n", str)
}

func TestStyle256(t *testing.T) {
	forceOpenColorRender()
	defer resetColorRender()

	at := assert.New(t)
	// empty
	s := S256()
	at.Equal("", s.String())
	at.Equal("MSG", s.Sprint("MSG"))

	// only fg
	s = S256(132)
	at.Equal("38;5;132", s.String())
	at.Equal("\x1b[38;5;132mMSG\x1b[0m", s.Sprint("MSG"))
	at.Equal("\x1b[38;5;132mMSG\x1b[0m", s.Sprintf("%s", "MSG"))

	// only bg
	s = S256(132)
	at.Equal("38;5;132", s.String())
	at.Equal("\x1b[38;5;132mMSG\x1b[0m", s.Sprint("MSG"))

	// fg and bg
	s = S256(132, 23)
	at.Equal("38;5;132;48;5;23", s.String())
	at.Equal("\x1b[38;5;132;48;5;23mMSG\x1b[0m", s.Sprint("MSG"))
	s = S256().Set(132, 23)
	at.Equal("38;5;132;48;5;23", s.String())
	at.Equal("\x1b[38;5;132;48;5;23mMSG\x1b[0m", s.Sprint("MSG"))
	s = S256().SetFg(132).SetBg(23)
	at.Equal("38;5;132;48;5;23", s.String())
	at.Equal("\x1b[38;5;132;48;5;23mMSG\x1b[0m", s.Sprint("MSG"))

	s = S256(132)

	// Color256.Print
	rewriteStdout()
	s.Print("MSG")
	str := restoreStdout()
	at.Equal("\x1b[38;5;132mMSG\x1b[0m", str)

	// Color256.Printf
	rewriteStdout()
	s.Printf("A %s", "MSG")
	str = restoreStdout()
	at.Equal("\x1b[38;5;132mA MSG\x1b[0m", str)

	// Color256.Println
	rewriteStdout()
	s.Println("MSG")
	str = restoreStdout()
	at.Equal("\x1b[38;5;132mMSG\x1b[0m\n", str)
}

/*************************************************************
 * test rgb color
 *************************************************************/

func TestRGBColor(t *testing.T) {
	forceOpenColorRender()
	defer resetColorRender()
	at := assert.New(t)

	// empty
	c := RGBColor{3: 99}
	at.True(c.IsEmpty())
	at.Equal("", c.String())

	// bg
	c = RGB(204, 204, 204, true)
	at.False(c.IsEmpty())
	at.Equal("48;2;204;204;204", c.String())

	// fg
	c = RGB(204, 204, 204)
	at.False(c.IsEmpty())
	at.Equal("38;2;204;204;204", c.String())

	// RGBColor.Sprint
	str := c.Sprint("msg")
	at.Equal("\x1b[38;2;204;204;204mmsg\x1b[0m", str)

	// RGBColor.Sprintf
	str = c.Sprintf("msg")
	at.Equal("\x1b[38;2;204;204;204mmsg\x1b[0m", str)
	at.Equal("[204 204 204]", fmt.Sprint(c.Values()))

	// RGBColor.Print
	rewriteStdout()
	c.Print("msg")
	str = restoreStdout()
	at.Equal("\x1b[38;2;204;204;204mmsg\x1b[0m", str)

	// RGBColor.Printf
	rewriteStdout()
	c.Printf("m%s", "sg")
	str = restoreStdout()
	at.Equal("\x1b[38;2;204;204;204mmsg\x1b[0m", str)

	// RGBColor.Println
	rewriteStdout()
	c.Println("msg")
	str = restoreStdout()
	at.Equal("\x1b[38;2;204;204;204mmsg\x1b[0m\n", str)
}

func TestRGBFromString(t *testing.T) {
	forceOpenColorRender()
	defer resetColorRender()
	at := assert.New(t)

	c := RGBFromString("170,187,204")
	at.Equal("\x1b[38;2;170;187;204mmsg\x1b[0m", c.Sprint("msg"))

	c = RGBFromString("170,187,204", true)
	at.Equal("\x1b[48;2;170;187;204mmsg\x1b[0m", c.Sprint("msg"))

	c = RGBFromString("170,187,")
	at.Equal("msg", c.Sprint("msg"))

	c = RGBFromString("")
	at.Equal("msg", c.Sprint("msg"))

	c = RGBFromString("170,187,error")
	at.Equal("msg", c.Sprint("msg"))
}

func TestHexToRGB(t *testing.T) {
	at := assert.New(t)
	c := HEX("ccc") // rgb: [204 204 204]
	at.False(c.IsEmpty())
	at.Equal("38;2;204;204;204", c.String())

	c = HEX("aabbcc") // rgb: [170 187 204]
	at.Equal("38;2;170;187;204", c.String())

	c = HEX("#aabbcc") // rgb: [170 187 204]
	at.Equal("38;2;170;187;204", c.String())

	c = HEX("0xad99c0") // rgb: [170 187 204]
	at.Equal("38;2;173;153;192", c.String())

	c = HEX(" ")
	at.True(c.IsEmpty())
	at.Equal("", c.String())

	c = HEX("!#$bbcc")
	at.Equal("", c.String())

	c = HEX("#invalid")
	at.Equal("", c.String())

	c = HEX("invalid code")
	at.Equal("", c.String())
}

func TestRGBStyle(t *testing.T) {
	forceOpenColorRender()
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
	rewriteStdout()
	s.Print("msg")
	str := restoreStdout()
	at.Equal("\x1b[38;2;20;144;234;48;2;234;78;23mmsg\x1b[0m", str)

	// RGBColor.Printf
	rewriteStdout()
	s.Printf("m%s", "sg")
	str = restoreStdout()
	at.Equal("\x1b[38;2;20;144;234;48;2;234;78;23mmsg\x1b[0m", str)

	// RGBColor.Println
	rewriteStdout()
	s.Println("msg")
	str = restoreStdout()
	at.Equal("\x1b[38;2;20;144;234;48;2;234;78;23mmsg\x1b[0m\n", str)
}

func TestOther(t *testing.T) {
	at := assert.New(t)

	at.True(IsConsole(os.Stdout))
	at.False(IsConsole(&bytes.Buffer{}))

	// IsMSys
	oldVal := os.Getenv("MSYSTEM")
	at.NoError(os.Setenv("MSYSTEM", "MINGW64"))
	at.True(IsMSys())
	at.NoError(os.Unsetenv("MSYSTEM"))
	at.False(IsMSys())
	_ = os.Setenv("MSYSTEM", oldVal)

	// TERM
	oldVal = os.Getenv("TERM")
	_ = os.Unsetenv("TERM")
	at.False(IsSupport256Color())

	at.NoError(os.Setenv("TERM", "xterm-vt220"))
	at.True(IsSupportColor())
	// revert
	if oldVal != "" {
		at.NoError(os.Setenv("TERM", oldVal))
	} else {
		at.NoError(os.Unsetenv("TERM"))
	}

	// ConEmuANSI
	oldVal = os.Getenv("ConEmuANSI")
	at.NoError(os.Setenv("ConEmuANSI", "ON"))
	at.True(IsSupportColor())
	// revert
	if oldVal != "" {
		at.NoError(os.Setenv("ConEmuANSI", oldVal))
	} else {
		at.NoError(os.Unsetenv("ConEmuANSI"))
	}

	// ANSICON
	oldVal = os.Getenv("ANSICON")
	at.NoError(os.Setenv("ANSICON", "189x2000 (189x43)"))
	at.True(IsSupportColor())
	// revert
	if oldVal != "" {
		at.NoError(os.Setenv("ANSICON", oldVal))
	} else {
		at.NoError(os.Unsetenv("ANSICON"))
	}
}

/*************************************************************
 * test helpers
 *************************************************************/

var oldVal bool

// force open color render for testing
func forceOpenColorRender() {
	oldVal = isSupportColor
	isSupportColor = true
}

func resetColorRender() {
	isSupportColor = oldVal
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
