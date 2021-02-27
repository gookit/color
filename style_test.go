package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStyle(t *testing.T) {
	// force open color render for testing
	buf := forceOpenColorRender()
	defer resetColorRender()

	is := assert.New(t)

	// IsEmpty
	s := Style{}
	is.True(s.IsEmpty())
	is.Equal("", s.String())

	is.Equal("97;40", Light.Code())
	is.Equal("97;40", Light.String())
	str := Light.Render("msg")
	is.Contains(str, "97")

	str = Danger.Sprint("msg")
	is.Contains(str, FgRed.String())

	str = Question.Render("msg")
	is.Contains(str, FgMagenta.String())

	str = Question.Render("msg", "More")
	is.Contains(str, FgMagenta.String())
	is.Contains(str, "msgMore")

	str = Question.Renderln("msg", "More")
	is.Contains(str, FgMagenta.String())
	is.Contains(str, "msg More")

	str = Secondary.Sprintf("m%s", "sg")
	is.Contains(str, FgDarkGray.String())

	// Style.Print
	Info.Print("MSG")
	is.Equal("\x1b[0;32mMSG\x1b[0m", buf.String())
	buf.Reset()

	// Style.Printf
	Info.Printf("A %s", "MSG")
	is.Equal("\x1b[0;32mA MSG\x1b[0m", buf.String())
	buf.Reset()

	// Style.Println
	Info.Println("MSG")
	is.Equal("\x1b[0;32mMSG\x1b[0m\n", buf.String())
	buf.Reset()

	Info.Println("MSG", "OK")
	is.Equal("\x1b[0;32mMSG OK\x1b[0m\n", buf.String())
	buf.Reset()

	s = GetStyle("err")
	is.False(s.IsEmpty())

	if isLikeInCmd {
		s.Print("msg")
		s.Printf("M%s", "sg")
		s.Println("Msg")
		is.Equal("\x1b[97;41mmsg\x1b[0m\x1b[97;41mMsg\x1b[0m\x1b[97;41mMsg\x1b[0m\n", buf.String())
		buf.Reset()
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
	buf := forceOpenColorRender()
	defer resetColorRender()

	is := assert.New(t)

	// Theme.Tips
	Info.Tips("MSG")
	is.Equal("\x1b[0;32mINFO: \x1b[0mMSG\n", buf.String())
	buf.Reset()

	// Theme.Prompt
	Info.Prompt("MSG")
	is.Equal("\x1b[0;32mINFO: MSG\x1b[0m\n", buf.String())
	buf.Reset()

	// Theme.Block
	Info.Block("MSG")
	is.Equal("\x1b[0;32mINFO:\n MSG\x1b[0m\n", buf.String())
	buf.Reset()

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

func TestStyleFunc(t *testing.T) {
	// force open color render for testing
	buf := forceOpenColorRender()
	defer resetColorRender()

	Infoln("color message")
	assert.Equal(t, "\x1b[0;32mcolor message\x1b[0m\n", buf.String())
	buf.Reset()

	Infof("color %s", "message")
	assert.Equal(t, "\x1b[0;32mcolor message\x1b[0m", buf.String())
	buf.Reset()

	Warnln("color message")
	assert.Equal(t, "\x1b[1;33mcolor message\x1b[0m\n", buf.String())
	buf.Reset()

	Warnf("color %s", "message")
	assert.Equal(t, "\x1b[1;33mcolor message\x1b[0m", buf.String())
	buf.Reset()

	Errorln("color message")
	assert.Equal(t, "\x1b[97;41mcolor message\x1b[0m\n", buf.String())
	buf.Reset()

	Errorf("color %s", "message")
	assert.Equal(t, "\x1b[97;41mcolor message\x1b[0m", buf.String())
	buf.Reset()
}

func TestSimplePrinter_Print(t *testing.T) {
	sp := &SimplePrinter{}
	sp.Printf("simple %s", "printer")
	sp.Infof("simple %s", "printer")
	sp.Warnf("simple %s", "printer")
	sp.Errorf("simple %s", "printer")
	sp.Print("simple printer\n")
	sp.Println("simple printer")
	sp.Infoln("simple printer")
	sp.Warnln("simple printer")
	sp.Errorln("simple printer")
}

func TestNewScheme(t *testing.T) {
	cs := NewDefaultScheme("test")

	cs.Infof("color %s\n", "scheme")
	cs.Warnf("color %s\n", "scheme")
	cs.Errorf("color %s\n", "scheme")
	cs.Infoln("color scheme")
	cs.Warnln("color scheme")
	cs.Errorln("color scheme")
	cs.Style("info").Println("color scheme")
}
