package color

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

var oldVal bool

// force open color render for testing
func forceOpenColorRender() {
	oldVal = isSupportColor
	isSupportColor = true
}

func resetColorRender() {
	isSupportColor = oldVal
}

func TestColor_Render(t *testing.T) {
	forceOpenColorRender()
	defer resetColorRender()
	at := assert.New(t)

	r := Bold.Render("text")
	at.Equal("\x1b[1mtext\x1b[0m", r)
	r = Bold.Text("text")
	at.Equal("\x1b[1mtext\x1b[0m", r)
	r = Bold.Sprint("text")
	at.Equal("\x1b[1mtext\x1b[0m", r)
}

func TestRenderCode(t *testing.T) {
	// force open color render for testing
	forceOpenColorRender()
	defer resetColorRender()

	at := assert.New(t)
	str := RenderCode("36;1", "Text")
	at.Contains(str, "\x1b[36;1m")
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

	str := c.Sprint("msg")
	at.Equal("\x1b[38;5;132mmsg\x1b[0m", str)
	str = c.Sprintf("msg")
	at.Equal("\x1b[38;5;132mmsg\x1b[0m", str)

	// bg
	c = Bit8(132, true)
	at.False(c.IsEmpty())
	at.Equal("48;5;132", c.String())
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
}

func TestRGBColor(t *testing.T) {
	forceOpenColorRender()
	defer resetColorRender()

	at := assert.New(t)

	// empty
	c := RGBColor{3: 99}
	at.True(c.IsEmpty())
	at.Equal(ResetCode, c.String())

	// fg
	c = RGB(204, 204, 204)
	at.False(c.IsEmpty())
	at.Equal("38;2;204;204;204", c.String())

	str := c.Sprint("msg")
	at.Equal("\x1b[38;2;204;204;204mmsg\x1b[0m", str)
	str = c.Sprintf("msg")
	at.Equal("\x1b[38;2;204;204;204mmsg\x1b[0m", str)

	// bg
	c = RGB(204, 204, 204, true)
	at.False(c.IsEmpty())
	at.Equal("48;2;204;204;204", c.String())
}

func TestHexToRGB(t *testing.T) {
	at := assert.New(t)
	rgb := HEX("ccc") // rgb: [204 204 204]
	at.Equal("38;2;204;204;204", rgb.String())

	rgb = HEX("aabbcc") // rgb: [170 187 204]
	at.Equal("38;2;170;187;204", rgb.String())

	rgb = HEX("0xad99c0") // rgb: [170 187 204]
	at.Equal("38;2;173;153;192", rgb.String())

	rgb = HEX("invalid code")
	at.Equal(ResetCode, rgb.String())
}
