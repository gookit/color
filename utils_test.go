package color

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtilFuncs(t *testing.T) {
	is := assert.New(t)

	// IsConsole
	is.True(IsConsole(os.Stdin))
	is.True(IsConsole(os.Stdout))
	is.True(IsConsole(os.Stderr))
	is.False(IsConsole(&bytes.Buffer{}))
	ff, err := os.OpenFile("README.md", os.O_WRONLY, 0)
	is.NoError(err)
	is.False(IsConsole(ff))

	// IsMSys
	oldVal := os.Getenv("MSYSTEM")
	is.NoError(os.Setenv("MSYSTEM", "MINGW64"))
	is.True(IsMSys())
	is.NoError(os.Unsetenv("MSYSTEM"))
	is.False(IsMSys())
	_ = os.Setenv("MSYSTEM", oldVal)

	is.NotEmpty(TermColorLevel())
	// is.NotEmpty(SupColorMark())
}

func TestDetectColorLevel(t *testing.T) {
	is := assert.New(t)

	// "COLORTERM=truecolor"
	mockOsEnvByText("COLORTERM=truecolor", func() {
		is.True(IsSupportColor())
		is.Equal(LevelRgb, DetectColorLevel())
		is.True(IsSupportRGBColor())
		is.True(IsSupportTrueColor())
	})

	// "FORCE_COLOR=on"
	mockOsEnvByText("FORCE_COLOR=on", func() {
		is.True(IsSupportColor())
		is.Equal(Level16, DetectColorLevel())
		is.False(IsSupportRGBColor())
		is.False(IsSupportTrueColor())
	})

	// TERMINAL_EMULATOR=JetBrains-JediTerm
	mockOsEnvByText(`
TERM=xterm-256color
TERMINAL_EMULATOR=JetBrains-JediTerm
ZSH_TMUX_TERM=screen-256color
`, func() {
		is.Equal(LevelRgb, DetectColorLevel())
		is.True(IsSupportRGBColor())
		is.True(IsSupport16Color())
		is.True(IsSupportColor())
	})
}

func TestIsDetectColorLevel_unix(t *testing.T) {
	if IsWindows() {
		return
	}
	is := assert.New(t)

	// no TERM env
	mockOsEnvByText("NO=none", func() {
		is.Equal(LevelNo, DetectColorLevel())
		is.False(IsSupportTrueColor())
		is.False(IsSupport256Color())
		is.False(IsSupportColor())
	})

	mockOsEnvByText("TERM=not-exist-value", func() {
		is.Equal(Level16, DetectColorLevel())
		is.False(IsSupportTrueColor())
		is.False(IsSupport256Color())
		is.True(IsSupportColor())
	})

	mockOsEnvByText("TERM=xterm", func() {
		is.Equal(Level256, DetectColorLevel())
		is.False(IsSupportTrueColor())
		is.True(IsSupport256Color())
		is.True(IsSupportColor())
	})

	mockOsEnvByText("TERM=screen-256color", func() {
		is.Equal(Level256, DetectColorLevel())
		is.False(IsSupportTrueColor())
		is.True(IsSupport256Color())
		is.True(IsSupportColor())
	})

	mockOsEnvByText("TERM=not-exist-256color", func() {
		is.Equal(Level256, DetectColorLevel())
		is.False(IsSupportTrueColor())
		is.True(IsSupport256Color())
		is.True(IsSupportColor())
	})

	// TERM_PROGRAM=Terminus
	mockOsEnvByText(`
TERMINUS_PLUGINS=
TERM=xterm-256color
TERM_PROGRAM=Terminus
ZSH_TMUX_TERM=screen-256color
`, func() {
		is.Equal(LevelRgb, DetectColorLevel())
		is.True(IsSupportRGBColor())
		is.True(IsSupport16Color())
		is.True(IsSupportColor())
	})

	// -------- tests on macOS ---------

	// TERM_PROGRAM=Apple_Terminal
	mockOsEnvByText(`
TERM_PROGRAM=Apple_Terminal
TERM=xterm-256color
TERM_PROGRAM_VERSION=433
TERM_SESSION_ID=F17907FE-DCA5-488D-829B-7AFA8B323753
ZSH_TMUX_TERM=screen-256color
`, func() {
		// fmt.Println(os.Environ())
		is.Equal(Level256, DetectColorLevel())
		is.False(IsSupportTrueColor())
		is.True(IsSupport256Color())
		is.True(IsSupport16Color())
		is.True(IsSupportColor())
	})

	// TERM_PROGRAM=iTerm.app
	mockOsEnvByText(`
ITERM_PROFILE=Default
TERM_PROGRAM_VERSION=3.4.5beta1
TERM_PROGRAM=iTerm.app
LC_TERMINAL=iTerm2
TERM=xterm-256color
ZSH_TMUX_TERM=screen-256color
`, func() {
		is.Equal(LevelRgb, DetectColorLevel())
		is.True(IsSupportTrueColor())
		is.True(IsSupport256Color())
		is.True(IsSupport16Color())
		is.True(IsSupportColor())
	})

	// TERM_PROGRAM=iTerm.app invalid version
	mockOsEnvByText(`
ITERM_PROFILE=Default
TERM_PROGRAM_VERSION=xx.beta
TERM_PROGRAM=iTerm.app
LC_TERMINAL=iTerm2
TERM=xterm-256color
ZSH_TMUX_TERM=screen-256color
`, func() {
		is.Equal(Level256, DetectColorLevel())
		is.False(IsSupportTrueColor())
		is.True(IsSupport256Color())
		is.True(IsSupport16Color())
		is.True(IsSupportColor())
	})

	// TERM_PROGRAM=iTerm.app no version env
	mockOsEnvByText(`
ITERM_PROFILE=Default
TERM_PROGRAM=iTerm.app
LC_TERMINAL=iTerm2
TERM=xterm-256color
ZSH_TMUX_TERM=screen-256color
`, func() {
		is.Equal(Level256, DetectColorLevel())
		is.False(IsSupportTrueColor())
		is.True(IsSupport256Color())
		is.True(IsSupport16Color())
		is.True(IsSupportColor())
	})

	// -------- tests on linux ---------
}

func TestIsDetectColorLevel_screen(t *testing.T) {
	if IsWindows() {
		return
	}
	is := assert.New(t)

	// COLORTERM=truecolor
	mockOsEnvByText(`
TERM=screen
COLORTERM=truecolor
`, func() {
		is.Equal(Level256, DetectColorLevel())
		is.False(IsSupportRGBColor())
		is.False(IsSupportTrueColor())
		is.True(IsSupport256Color())
		is.True(IsSupportColor())
	})

	// TERM_PROGRAM=Apple_Terminal use screen
	mockOsEnvByText(`
TERM_PROGRAM=Apple_Terminal
TERM=screen
TERM_PROGRAM_VERSION=433
TERM_SESSION_ID=F17907FE-DCA5-488D-829B-7AFA8B323753
ZSH_TMUX_TERM=screen-256color
`, func() {
		// fmt.Println(os.Environ())
		is.Equal(Level256, DetectColorLevel())
		is.False(IsSupportTrueColor())
		is.True(IsSupport256Color())
		is.True(IsSupport16Color())
		is.True(IsSupportColor())
	})

	// TERM_PROGRAM=iTerm.app use screen
	mockOsEnvByText(`
TERM=screen
TERMCAP=SC|screen|VT 100/ANSI X3.64 virtual terminal:\
LC_TERMINAL_VERSION=3.4.5beta1
ITERM_PROFILE=Default
TERM_PROGRAM_VERSION=3.4.5beta1
TERM_PROGRAM=iTerm.app
LC_TERMINAL=iTerm2
ZSH_TMUX_TERM=screen
`, func() {
		is.Equal(Level256, DetectColorLevel())
		is.False(IsSupportTrueColor())
		is.True(IsSupport256Color())
		is.True(IsSupport16Color())
		is.True(IsSupportColor())
	})

	// TERM_PROGRAM=Terminus use screen
	mockOsEnvByText(`
TERM=screen
TERMCAP=SC|screen|VT 100/ANSI X3.64 virtual terminal:\
TERMINUS_PLUGINS=
TERM_PROGRAM=Terminus
ZSH_TMUX_TERM=screen
`, func() {
		is.Equal(Level256, DetectColorLevel())
		is.False(IsSupportTrueColor())
		is.True(IsSupport256Color())
		is.True(IsSupport16Color())
		is.True(IsSupportColor())
	})

	// TERMINAL_EMULATOR=JetBrains-JediTerm use screen
	mockOsEnvByText(`
TERM=screen
TERMCAP=SC|screen|VT 100/ANSI X3.64 virtual terminal:\
TERMINAL_EMULATOR=JetBrains-JediTerm
ZSH_TMUX_TERM=screen
`, func() {
		is.Equal(Level256, DetectColorLevel())
		is.False(IsSupportTrueColor())
		is.True(IsSupport256Color())
		is.True(IsSupport16Color())
		is.True(IsSupportColor())
	})
}

func TestIsDetectColorLevel_win(t *testing.T) {
	if !IsWindows() {
		return
	}
	is := assert.New(t)

	// ConEmuANSI
	mockEnvValue("ConEmuANSI", "ON", func(_ string) {
		is.Equal(LevelRgb, DetectColorLevel())
		is.True(IsSupportColor())
		is.True(IsSupport256Color())
		is.True(IsSupportTrueColor())
	})

	// WSL_DISTRO_NAME=Debian
	mockEnvValue("WSL_DISTRO_NAME", "Debian", func(_ string) {
		is.True(IsSupportColor())
	})

	// ANSICON
	mockEnvValue("ANSICON", "189x2000 (189x43)", func(_ string) {
		is.True(IsSupportColor())
		// is.Equal(Level256, DetectColorLevel())
		// is.Equal("TERM=xterm-256color", SupColorMark())
	})
}

func TestRgbTo256Table(t *testing.T) {
	index := 0
	for hex, c256 := range RgbTo256Table() {
		Hex(hex).Print("RGB:", hex)
		fmt.Print(" = ")
		C256(c256).Print("C256:", c256)
		fmt.Print(" | ")
		index++
		if index%5 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()

	assert.Equal(t, uint8(0x92), RgbTo256(170, 187, 204))
}

func TestC256ToRgbV1(t *testing.T) {
	for i := 0; i < 256; i++ {
		c256 := uint8(i)
		C256(c256).Printf("C256:%d", c256)
		fmt.Print(" => ")
		rgb := C256ToRgbV1(c256)
		RGBFromSlice(rgb).Printf("RGB:%v | ", rgb)
		// assert.Equal(t, item.want, rgb, fmt.Sprint("256 code:", c256))
		if i%4 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func TestC256ToRgb(t *testing.T) {
	for i := 0; i < 256; i++ {
		c256 := uint8(i)
		C256(c256).Printf("C256:%d", c256)
		fmt.Print(" => ")
		rgb := C256ToRgb(c256)
		RGBFromSlice(rgb).Printf("RGB:%v | ", rgb)
		// assert.Equal(t, item.want, rgb, fmt.Sprint("256 code:", c256))
		if i%4 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}

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

func TestRgb2basic(t *testing.T) {
	assert.Equal(t, uint8(31), Rgb2basic(134, 56, 56, false))
	assert.Equal(t, uint8(41), Rgb2basic(134, 56, 56, true))
	assert.Equal(t, uint8(46), Rgb2basic(57, 187, 226, true))
}
