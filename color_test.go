package color

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
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

	// set a style tag
	Tag("info").Println("info style text")

	// use info style tips
	Tips("info").Print("tips style text")

	// use info style blocked tips
	LiteTips("info").Print("blocked tips style text")
}

/*************************************************************
 * test global methods
 *************************************************************/

func TestSet(t *testing.T) {
	at := assert.New(t)

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
	Set(FgGreen)
	str := restoreStdout()
	at.Equal("\x1b[32m", str)
	// unset
	rewriteStdout()
	Reset()
	str = restoreStdout()
	at.Equal("\x1b[0m", str)
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
 * test 16 color
 *************************************************************/

func TestColor16(t *testing.T) {
	forceOpenColorRender()
	defer resetColorRender()
	at := assert.New(t)

	at.True(Bold.IsValid())
	r := Bold.Render("text")
	at.Equal("\x1b[1mtext\x1b[0m", r)
	r = Bold.Text("text")
	at.Equal("\x1b[1mtext\x1b[0m", r)
	r = Bold.Sprint("text")
	at.Equal("\x1b[1mtext\x1b[0m", r)

	str := Red.Sprintf("A %s", "MSG")
	at.Equal("\x1b[31mA MSG\x1b[0m", str)

	// Color.Print
	rewriteStdout()
	FgGray.Print("MSG")
	str = restoreStdout()
	at.Equal("\x1b[90mMSG\x1b[0m", str)

	// Color.Printf
	rewriteStdout()
	BgGray.Printf("A %s", "MSG")
	str = restoreStdout()
	at.Equal("\x1b[100mA MSG\x1b[0m", str)

	// Color.Println
	rewriteStdout()
	BgGray.Println("MSG")
	str = restoreStdout()
	at.Equal("\x1b[100mMSG\x1b[0m\n", str)

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
	at.Equal(ResetCode, c.String())

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
	at.Equal(ResetCode, c.String())

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
}

func TestHexToRGB(t *testing.T) {
	at := assert.New(t)
	rgb := HEX("ccc") // rgb: [204 204 204]
	at.False(rgb.IsEmpty())
	at.Equal("38;2;204;204;204", rgb.String())

	rgb = HEX("aabbcc") // rgb: [170 187 204]
	at.Equal("38;2;170;187;204", rgb.String())

	rgb = HEX("#aabbcc") // rgb: [170 187 204]
	at.Equal("38;2;170;187;204", rgb.String())

	rgb = HEX("0xad99c0") // rgb: [170 187 204]
	at.Equal("38;2;173;153;192", rgb.String())

	rgb = HEX(" ")
	at.True(rgb.IsEmpty())
	at.Equal(ResetCode, rgb.String())

	rgb = HEX("!#$bbcc")
	at.Equal(ResetCode, rgb.String())

	rgb = HEX("#invalid")
	at.Equal(ResetCode, rgb.String())

	rgb = HEX("invalid code")
	at.Equal(ResetCode, rgb.String())
}

func TestRGBStyle(t *testing.T) {
	forceOpenColorRender()
	defer resetColorRender()
	at := assert.New(t)

	s := &RGBStyle{}
	at.True(s.IsEmpty())
	// NewRGBStyle
	s = NewRGBStyle(RGB(20, 144, 234), RGB(234, 78, 23))
	at.False(s.IsEmpty())
	// HEXStyle
	s = HEXStyle("555", "eee")
	at.False(s.IsEmpty())
	// RGBStyleFromString
	s = RGBStyleFromString("20, 144, 234", "234, 78, 23")
	at.False(s.IsEmpty())
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

// usage:
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
	os.Stdout.Close()
	// restore
	os.Stdout = oldStdout
	oldStdout = nil

	// read data
	out, _ := ioutil.ReadAll(newReader)

	// close reader
	newReader.Close()
	newReader = nil

	return string(out)
}
