package color

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStyles(t *testing.T) {
	// force open color render for testing
	forceOpenColorRender()
	defer resetColorRender()

	at := assert.New(t)

	// IsEmpty
	s := Style{}
	at.True(s.IsEmpty())

	at.Equal("97", Light.String())
	str := Light.Render("msg")
	at.Contains(str, "97")

	str = Danger.Render("msg")
	at.Contains(str, FgRed.String())

	str = Question.Render("msg")
	at.Contains(str, FgMagenta.String())

	// add new
	AddStyle("new0", Style{OpFastBlink})
	s = GetStyle("new0")
	at.False(s.IsEmpty())
	delete(Styles, "new0")
}

func TestThemes(t *testing.T) {
	at := assert.New(t)

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
}
