package color

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplaceTag(t *testing.T) {
	// force open color render for testing
	forceOpenColorRender()
	defer resetColorRender()

	at := assert.New(t)

	// sample 1
	r := String("<err>text</>")
	at.NotContains(r, "<")
	at.NotContains(r, ">")

	// disable color
	Enable = false
	r = Text("<err>text</>")
	at.Equal("text", r)
	Enable = true

	// sample 2
	s := "abc <err>err-text</> def <info>info text</>"
	r = ReplaceTag(s)
	at.NotContains(r, "<")
	at.NotContains(r, ">")

	// sample 3
	s = `abc <err>err-text</> 
def <info>info text
</>`
	r = ReplaceTag(s)
	at.NotContains(r, "<")
	at.NotContains(r, ">")

	// sample 4
	s = "abc <err>err-text</> def <err>err-text</> "
	r = ReplaceTag(s)
	at.NotContains(r, "<")
	at.NotContains(r, ">")

	// sample 5
	s = "abc <err>err-text</> def <d>"
	r = ReplaceTag(s)
	at.NotContains(r, "<err>")
	at.Contains(r, "<d>")

	// sample 6
	s = "custom tag: <fg=yellow;bg=black;op=underscore;>hello, welcome</>"
	r = ReplaceTag(s)
	at.NotContains(r, "<")
	at.NotContains(r, ">")
}

func TestPrint(t *testing.T) {
	// force open color render for testing
	forceOpenColorRender()
	defer resetColorRender()
	at := assert.New(t)

	at.True(len(GetColorTags()) > 0)
	at.True(IsDefinedTag("info"))
	at.Equal("0;32", GetTagCode("info"))
	at.Equal("", GetTagCode("not-exist"))

	s := Sprint("<red>MSG</>")
	at.Equal("\x1b[0;31mMSG\x1b[0m", s)

	s = Sprint("<red>H</><green>I</>")
	at.Equal("\x1b[0;31mH\x1b[0m\x1b[0;32mI\x1b[0m", s)

	s = Sprintf("<red>%s</>", "MSG")
	at.Equal("\x1b[0;31mMSG\x1b[0m", s)

	// Print
	rewriteStdout()
	Print("<red>MSG</>")
	s = restoreStdout()
	at.Equal("\x1b[0;31mMSG\x1b[0m", s)

	// Printf
	rewriteStdout()
	Printf("<red>%s</>", "MSG")
	s = restoreStdout()
	at.Equal("\x1b[0;31mMSG\x1b[0m", s)

	// Println
	rewriteStdout()
	Println("<red>MSG</>")
	s = restoreStdout()
	at.Equal("\x1b[0;31mMSG\x1b[0m\n", s)
}

func TestWrapTag(t *testing.T) {
	at := assert.New(t)
	at.Equal("<info>text</>", WrapTag("text", "info"))
}

func TestApplyTag(t *testing.T) {
	forceOpenColorRender()
	defer resetColorRender()
	at := assert.New(t)
	at.Equal("\x1b[0;32mMSG\x1b[0m", ApplyTag("info", "MSG"))
}

func TestClearTag(t *testing.T) {
	at := assert.New(t)
	at.Equal("text", ClearTag("text"))
	at.Equal("text", ClearTag("<err>text</>"))
	at.Equal("abc error def info text", ClearTag("abc <err>error</> def <info>info text</>"))

	str := `abc <err>err-text</> 
def <info>info text
</>`
	ret := ClearTag(str)
	at.Contains(ret, "def info")
	at.NotContains(ret, "</>")

	str = "abc <err>text</> def<d>"
	ret = ClearTag(str)
	at.Equal("abc text def", ret)
	at.NotContains(ret, "<err>")
}
