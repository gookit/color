package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStyle(t *testing.T) {
	// force open color render for testing
	forceOpenColorRender()
	defer resetColorRender()

	is := assert.New(t)

	// IsEmpty
	s := Style{}
	is.True(s.IsEmpty())
	is.Equal("", s.String())

	is.Equal("97;40", Light.String())
	str := Light.Render("msg")
	is.Contains(str, "97")

	str = Danger.Sprint("msg")
	is.Contains(str, FgRed.String())

	str = Question.Render("msg")
	is.Contains(str, FgMagenta.String())

	str = Secondary.Sprintf("m%s", "sg")
	is.Contains(str, FgDarkGray.String())

	// Style.Print
	rewriteStdout()
	Info.Print("MSG")
	str = restoreStdout()
	if isLikeInCmd {
		is.Equal("MSG", str)
	} else {
		is.Equal("\x1b[0;32mMSG\x1b[0m", str)
	}

	// Style.Printf
	rewriteStdout()
	Info.Printf("A %s", "MSG")
	str = restoreStdout()
	if isLikeInCmd {
		is.Equal("A MSG", str)
	} else {
		is.Equal("\x1b[0;32mA MSG\x1b[0m", str)
	}

	// Style.Println
	rewriteStdout()
	Info.Println("MSG")
	str = restoreStdout()
	if isLikeInCmd {
		is.Equal("MSG\n", str)
	} else {
		is.Equal("\x1b[0;32mMSG\x1b[0m\n", str)
	}

	rewriteStdout()
	Info.Println("MSG", "OK")
	str = restoreStdout()
	if isLikeInCmd {
		is.Equal("MSG OK\n", str)
	} else {
		is.Equal("\x1b[0;32mMSG OK\x1b[0m\n", str)
	}

	s = GetStyle("err")
	is.False(s.IsEmpty())

	if isLikeInCmd {
		rewriteStdout()
		s.Print("msg")
		s.Printf("m%s", "sg")
		s.Println("msg")
		str = restoreStdout()
		is.Equal("msgmsgmsg\n", str)
	}

	// add new
	s = GetStyle("new0")
	is.True(s.IsEmpty())
	AddStyle("new0", Style{OpFastBlink})
	s = GetStyle("new0")
	is.False(s.IsEmpty())
	delete(Styles, "new0")

	// add new
	s = GetStyle("new1")
	is.True(s.IsEmpty())
	New(OpStrikethrough).Save("new1")
	s = GetStyle("new1")
	is.False(s.IsEmpty())
	delete(Styles, "new1")
}

func TestThemes(t *testing.T) {
	// force open color render for testing
	forceOpenColorRender()
	defer resetColorRender()

	is := assert.New(t)

	// Theme.Tips
	rewriteStdout()
	Info.Tips("MSG")
	str := restoreStdout()
	if isLikeInCmd {
		is.Equal("INFO: MSG\n", str)
	} else {
		is.Equal("\x1b[0;32mINFO: \x1b[0mMSG\n", str)
	}

	// Theme.Prompt
	rewriteStdout()
	Info.Prompt("MSG")
	str = restoreStdout()
	if isLikeInCmd {
		is.Equal("INFO: MSG\n", str)
	} else {
		is.Equal("\x1b[0;32mINFO: MSG\x1b[0m\n", str)
	}

	// Theme.Block
	rewriteStdout()
	Info.Block("MSG")
	str = restoreStdout()
	if isLikeInCmd {
		is.Equal("INFO:\n MSG\n", str)
	} else {
		is.Equal("\x1b[0;32mINFO:\n MSG\x1b[0m\n", str)
	}

	theme := GetTheme("info")
	is.NotNil(theme)
	theme = GetTheme("not-exist")
	is.Nil(theme)

	// add new
	AddTheme("new0", Style{OpFastBlink})
	theme = GetTheme("new0")
	is.NotNil(theme)
	delete(Themes, "new0")
	theme = GetTheme("new0")
	is.Nil(theme)

	// add new
	theme = GetTheme("new1")
	is.Nil(theme)

	theme = NewTheme("new1", Style{OpFastBlink})
	theme.Save()
	theme = GetTheme("new1")
	is.NotNil(theme)

	delete(Themes, "new1")
	theme = GetTheme("new1")
	is.Nil(theme)
}
