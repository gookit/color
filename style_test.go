package color

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStyles(t *testing.T) {
	art := assert.New(t)

	str := Light.Render("msg")
	art.Contains(str, "107")

	str = Danger.Render("msg")
	art.Contains(str, FgRed.String())

	str = Question.Render("msg")
	art.Contains(str, FgMagenta.String())
}
