package color

import (
	"testing"

	"github.com/gookit/assert"
)

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
		t.Skip("skip on windows")
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

	mockOsEnvByText("WSL_DISTRO_NAME=Debian", func() {
		is.Equal(LevelNo, DetectColorLevel())
		is.False(IsSupportTrueColor())
		is.False(IsSupport256Color())
		is.False(IsSupportColor())
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
		t.Skip("skip test on non-windows")
		return
	}
	is := assert.New(t)
	// EnableDebug()
	// defer ResetDebug()

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

	// TERM=screen
	mockOsEnvByText(`
COLORTERM=24bit
TERM=screen
`, func() {
		lv := DetectColorLevel()
		is.Eq(Level256, lv)
	})

	mockOsEnvByText(`
FORCE_COLOR=ON
TERM=xterm-256color
`, func() {
		lv := DetectColorLevel()
		is.Eq(Level256, lv)
	})

	mockOsEnvByText(`
TERM_PROGRAM=Terminus
`, func() {
		lv := DetectColorLevel()
		is.Eq(LevelRgb, lv)
	})

	mockOsEnvByText(`
TERM_PROGRAM=Apple_Terminal
`, func() {
		lv := DetectColorLevel()
		is.Eq(Level256, lv)
	})

	mockOsEnvByText(`
TERM_PROGRAM=iTerm.app
TERM_PROGRAM_VERSION=4.5
`, func() {
		lv := DetectColorLevel()
		is.Eq(Level256, lv)
	})
}
