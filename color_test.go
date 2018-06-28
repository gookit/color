package color

import (
	"testing"
	c "github.com/smartystreets/goconvey/convey"
)

func TestRenderCodes(t *testing.T) {
	c.Convey("test RenderCodes", t, func() {
		c.So(RenderCodes("36;1", "Text"), c.ShouldContainSubstring, "36;1")
	})
}

func TestClearCode(t *testing.T) {
	c.Convey("test clear color codes", t, func() {
		s := "\033[36;1mText\x1b[0m"
		c.Convey("should = 'Text'", func() {
			c.So(ClearCode(s), c.ShouldEqual, "Text")
		})

		s1 := "\033[36;1mText\x1b[0m other"
		c.Convey("should = 'Text other'", func() {
			c.So(ClearCode(s1), c.ShouldEqual, "Text other")
		})
	})
}
