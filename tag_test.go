package color

import (
	"testing"
	c "github.com/smartystreets/goconvey/convey"
)

func TestWrapTag(t *testing.T) {
	c.Convey("test wrap a tag", t, func() {
		c.So(WrapTag("text", "info"), c.ShouldEqual, "<info>text</>")
	})
}

func TestReplaceTag(t *testing.T) {
	c.Convey("test parse color tags", t, func() {
		c.Convey("sample 1", func() {
			r := ReplaceTag("<err>text</>")
			c.So(r, c.ShouldNotContainSubstring, "<")
			c.So(r, c.ShouldNotContainSubstring, ">")
		})

		c.Convey("sample 2", func() {
			s := "abc <err>err-text</> def <info>info text</>"
			r := ReplaceTag(s)
			c.So(r, c.ShouldNotContainSubstring, "<")
			c.So(r, c.ShouldNotContainSubstring, ">")
		})

		c.Convey("sample 3", func() {
			s := `abc <err>err-text</> 
def <info>info text
</>`
			r := ReplaceTag(s)
			c.So(r, c.ShouldNotContainSubstring, "<")
			c.So(r, c.ShouldNotContainSubstring, ">")
		})

		c.Convey("sample 4", func() {
			s := "abc <err>err-text</> def <err>err-text</> "
			r := ReplaceTag(s)
			c.So(r, c.ShouldNotContainSubstring, "<")
			c.So(r, c.ShouldNotContainSubstring, ">")
		})

		c.Convey("sample 5", func() {
			s := "abc <err>err-text</> def <d>"
			r := ReplaceTag(s)
			c.So(r, c.ShouldNotContainSubstring, "<err>")
			c.So(r, c.ShouldContainSubstring, "<d>")
		})
	})
}

func TestClearTag(t *testing.T) {
	s1 := "<err>text</>"
	c.Convey(s1+" -> 'text'", t, func() {
		c.So(ClearTag(s1), c.ShouldEqual, "text")
	})

	s2 := "abc <err>error</> def <info>info text</>"

	c.Convey(s1+" -> 'abc error def info text'", t, func() {
		c.So(ClearTag(s2), c.ShouldEqual, "abc error def info text")
	})

	s3 := `abc <err>err-text</> 
def <info>info text
</>`
	c.Convey("clear color tags", t, func() {
		r := ClearTag(s3)

		c.So(r, c.ShouldNotContainSubstring, "</>")
		c.So(r, c.ShouldContainSubstring, "def in")

		c.Convey("sample 1", func() {
			s := "abc <err>text</> def<d>"
			r := ClearTag(s)
			c.So(r, c.ShouldNotContainSubstring, "<err>")
			c.So(r, c.ShouldEqual, "abc text def")
		})
	})
}
