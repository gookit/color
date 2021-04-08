package color

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xo/terminfo"
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

	fmt.Println("support color:", SupportColor())
	fmt.Println("test color.Set() on OS:", runtime.GOOS)

	// disable
	old := Disable()
	num, err := Set(FgGreen)
	is.Nil(err)
	is.Equal(0, num)

	num, err = Reset()
	is.Nil(err)
	is.Equal(0, num)
	Enable = old // revert

	// set enable
	Enable = true
	if os.Getenv("GITHUB_ACTION") != "" {
		fmt.Println("Skip run the tests on Github Action")
		return
	}

	num, err = Set(FgGreen)
	is.True(num > 0)
	is.NoError(err)
	_, err = Reset()
	is.NoError(err)

	if runtime.GOOS == "windows" {
		fd := uintptr(syscall.Stdout)
		// if run test by goland, will return false
		if IsTerminal(fd) {
			fmt.Println("- IsTerminal return TRUE")
		} else {
			fmt.Println("- IsTerminal return FALSE")
		}
	} else {
		is.False(IsLikeInCmd())
		is.Empty(InnerErrs())
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

func TestSupportColor(t *testing.T) {
	is := assert.New(t)

	if SupportTrueColor() {
		is.True(SupportColor())
		is.True(Support256Color())
	}

	if Support256Color() {
		is.True(SupportColor())
	} else {
		is.False(SupportTrueColor())
	}

	if false == SupportColor() {
		is.False(Support256Color())
		is.False(SupportTrueColor())
	}
}

func TestRenderCode(t *testing.T) {
	// force open color render for testing
	oldVal = ForceColor()
	defer resetColorRender()

	is := assert.New(t)
	is.True(SupportColor())

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

	is.Empty(InnerErrs())
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
	is := assert.New(t)

	is.True(Bold.IsValid())
	r := Bold.Render("text")
	is.Equal("\x1b[1mtext\x1b[0m", r)
	r = LightYellow.Text("text")
	is.Equal("\x1b[93mtext\x1b[0m", r)
	r = LightWhite.Sprint("text")
	is.Equal("\x1b[97mtext\x1b[0m", r)
	r = White.Render("test", "spaces")
	is.Equal("\x1b[37mtestspaces\x1b[0m", r)
	r = Black.Renderln("test", "spaces")
	is.Equal("\x1b[30mtest spaces\x1b[0m", r)

	str := Red.Sprintf("A %s", "MSG")
	is.Equal("red", Red.Name())
	is.Equal("\x1b[31mA MSG\x1b[0m", str)

	// Color.Print
	FgGray.Print("MSG")
	str = buf.String()
	is.Equal("\x1b[90mMSG\x1b[0m", str)
	buf.Reset()

	// Color.Printf
	BgGray.Printf("A %s", "MSG")
	str = buf.String()
	is.Equal("\x1b[100mA MSG\x1b[0m", str)
	buf.Reset()

	// Color.Println
	LightMagenta.Println("MSG")
	str = buf.String()
	is.Equal("\x1b[95mMSG\x1b[0m\n", str)
	is.Equal("lightMagenta", LightMagenta.Name())
	buf.Reset()

	LightMagenta.Println()
	str = buf.String()
	is.Equal("\n", str)
	buf.Reset()

	LightCyan.Print("msg")
	LightRed.Printf("m%s", "sg")
	LightGreen.Println("msg")
	str = buf.String()
	is.Equal("\x1b[96mmsg\x1b[0m\x1b[91mmsg\x1b[0m\x1b[92mmsg\x1b[0m\n", str)
	buf.Reset()

	// Color.Darken
	blue := LightBlue.Darken()
	is.Equal(94, int(LightBlue))
	is.Equal(34, int(blue))
	c := Color(120).Darken()
	is.Equal(120, int(c))

	// Color.Light
	lightCyan := Cyan.Light()
	is.Equal(36, int(Cyan))
	is.Equal(96, int(lightCyan))
	c = Bit4(120).Light()
	is.Equal(120, int(c))

	// Colors vars
	_, ok := FgColors["red"]
	is.True(ok)
	_, ok = ExFgColors["lightRed"]
	is.True(ok)
	_, ok = BgColors["red"]
	is.True(ok)
	_, ok = ExBgColors["lightRed"]
	is.True(ok)
}

func TestColor_RGB(t *testing.T) {
	fmt.Println("------- 16-color code:")
	for u, s := range Basic2name {
		if u < 10 { // is option
			continue
		}

		fmt.Println(s, Color(u).RGB().Hex(), Color(u).RGB().Color())
	}

	fmt.Println("------- 256-color code:")
	for u, s := range Basic2name {
		if u < 10 { // is option
			continue
		}

		fmt.Println(s, Color(u).RGB().Hex(), Color(u).RGB().C256().Value())
	}
}

func TestColor_C256(t *testing.T) {

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
	for name, c := range AllOptions {
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

func TestQuickFunc(t *testing.T) {
	// inline func
	testFuncs := []func(...interface{}) {
		Redp,
		Greenp,
	}
	for _, fn := range testFuncs {
		fn("inline message,")
	}
}

/*************************************************************
 * test 256 color
 *************************************************************/

func TestColor256(t *testing.T) {
	is := assert.New(t)
	c := Bit8(132)
	c.Print("c256 message")
	is.Equal(uint8(132), c.Value())
	is.Equal("38;5;132", c.String())
	rgb := c.RGBColor()
	rgb.Print(" => to rgb message\n")
	is.Equal([]int{175, 95, 135}, rgb.Values())
	is.Equal("38;2;175;95;135", rgb.String())

	buf := forceOpenColorRender()
	defer resetColorRender()

	// empty
	c = Color256{1: 99}
	is.False(c.IsBg())
	is.True(c.IsEmpty())
	is.Equal("", c.String())

	// fg
	c = Bit8(132)
	is.True(c.IsFg())
	is.False(c.IsBg())
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
	is := assert.New(t)

	ForceColor()
	s := S256(192, 38)
	s.Println("style 256 colored text")
	is.Equal("\x1b[38;5;192;48;5;38m MSG \x1b[0m", s.Sprint(" MSG "))

	s.SetOpts(Opts{OpUnderscore})
	s.Println("style 256 colored text - with option OpUnderscore")
	is.Equal("\x1b[38;5;192;48;5;38;4m MSG \x1b[0m", s.Sprint(" MSG "))

	s.AddOpts(OpBold)
	s.Println("style 256 colored text - add option OpBold")
	is.Equal("\x1b[38;5;192;48;5;38;4;1m MSG \x1b[0m", s.Sprint(" MSG "))

	buf := forceOpenColorRender()
	defer resetColorRender()

	// empty
	s = S256()
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
	s = S256().Set(132, 23, OpStrikethrough)
	is.Equal("38;5;132;48;5;23;9", s.String())

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

func TestOpts_Add(t *testing.T) {
	is := assert.New(t)

	op := Opts{OpBold, OpBlink}
	is.False(op.IsEmpty())
	is.Equal("1;5", op.String())

	op.Add(Color(45))
	is.Equal("1;5", op.String())

	op.Add(OpUnderscore)
	is.Equal("1;5;4", op.String())
}

/*************************************************************
 * test helpers
 *************************************************************/

var oldVal terminfo.ColorLevel

// force open color render for testing
func forceOpenColorRender() *bytes.Buffer {
	oldVal = colorLevel
	ForceOpenColor()

	// set output for test
	buf := new(bytes.Buffer)
	SetOutput(buf)

	return buf
}

func resetColorRender() {
	colorLevel = oldVal
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

func mockOsEnvByText(envText string, fn func()) {
	ss := strings.Split(envText, "\n")
	mp := make(map[string]string, len(ss))
	for _, line := range ss {
		if line = strings.TrimSpace(line); line == "" {
			continue
		}
		nodes := strings.SplitN(line, "=", 2)

		if len(nodes) < 2 {
			mp[nodes[0]] = ""
		} else {
			mp[nodes[0]] = nodes[1]
		}
	}

	mockOsEnv(mp, fn)
}

func mockOsEnv(mp map[string]string, fn func()) {
	envBak := os.Environ()

	os.Clearenv()
	for key, val := range mp {
		_ = os.Setenv(key, val)
	}

	fn()

	os.Clearenv()
	for _, str := range envBak {
		nodes := strings.SplitN(str, "=", 2)
		_ = os.Setenv(nodes[0], nodes[1])
	}
}
