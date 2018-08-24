package color

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStyle(t *testing.T) {
	// force open color render for testing
	forceOpenColorRender()
	defer resetColorRender()

	at := assert.New(t)

	// IsEmpty
	s := Style{}
	at.True(s.IsEmpty())
	at.Equal("", s.String())

	at.Equal("97;40", Light.String())
	str := Light.Render("msg")
	at.Contains(str, "97")

	str = Danger.Sprint("msg")
	at.Contains(str, FgRed.String())

	str = Question.Render("msg")
	at.Contains(str, FgMagenta.String())

	str = Secondary.Sprintf("m%s", "sg")
	at.Contains(str, FgDarkGray.String())

	// Style.Print
	rewriteStdout()
	Info.Print("MSG")
	str = restoreStdout()
	at.Equal("\x1b[0;32mMSG\x1b[0m", str)

	// Style.Printf
	rewriteStdout()
	Info.Printf("A %s", "MSG")
	str = restoreStdout()
	at.Equal("\x1b[0;32mA MSG\x1b[0m", str)

	// Style.Println
	rewriteStdout()
	Info.Println("MSG")
	str = restoreStdout()
	at.Equal("\x1b[0;32mMSG\x1b[0m\n", str)

	s = GetStyle("err")
	at.False(s.IsEmpty())

	old := isLikeInCmd
	isLikeInCmd = true
	rewriteStdout()
	s.Print("msg")
	s.Printf("m%s", "sg")
	s.Println("msg")
	str = restoreStdout()
	isLikeInCmd = old

	// add new
	s = GetStyle("new0")
	at.True(s.IsEmpty())
	AddStyle("new0", Style{OpFastBlink})
	s = GetStyle("new0")
	at.False(s.IsEmpty())
	delete(Styles, "new0")

	// add new
	s = GetStyle("new1")
	at.True(s.IsEmpty())
	New(OpStrikethrough).Save("new1")
	s = GetStyle("new1")
	at.False(s.IsEmpty())
	delete(Styles, "new1")
}

func TestThemes(t *testing.T) {
	// force open color render for testing
	forceOpenColorRender()
	defer resetColorRender()

	at := assert.New(t)

	// Theme.Tips
	rewriteStdout()
	Info.Tips("MSG")
	str := restoreStdout()
	at.Equal("\x1b[0;32mINFO: \x1b[0mMSG\n", str)

	// Theme.Prompt
	rewriteStdout()
	Info.Prompt("MSG")
	str = restoreStdout()
	at.Equal("\x1b[0;32mINFO: MSG\x1b[0m\n", str)

	// Theme.Block
	rewriteStdout()
	Info.Block("MSG")
	str = restoreStdout()
	at.Equal("\x1b[0;32mINFO:\n MSG\x1b[0m\n", str)

	theme := GetTheme("info")
	at.NotNil(theme)
	theme = GetTheme("not-exist")
	at.Nil(theme)

	// add new
	AddTheme("new0", Style{OpFastBlink})
	theme = GetTheme("new0")
	at.NotNil(theme)
	delete(Themes, "new0")
	theme = GetTheme("new0")
	at.Nil(theme)

	// add new
	theme = GetTheme("new1")
	at.Nil(theme)

	theme = NewTheme("new1", Style{OpFastBlink})
	theme.Save()
	theme = GetTheme("new1")
	at.NotNil(theme)

	delete(Themes, "new1")
	theme = GetTheme("new1")
	at.Nil(theme)
}
