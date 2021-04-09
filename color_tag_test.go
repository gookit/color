package color

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceTag(t *testing.T) {
	// force open color render for testing
	forceOpenColorRender()
	defer resetColorRender()

	is := assert.New(t)

	// sample 1
	r := String("<err>text</>")
	is.NotContains(r, "<")
	is.NotContains(r, ">")

	// disable color
	Enable = false
	r = Text("<err>text</>")
	is.Equal("text", r)
	Enable = true

	// sample 2
	s := "abc <err>err-text</> def <info>info text</>"
	r = ReplaceTag(s)
	is.NotContains(r, "<")
	is.NotContains(r, ">")

	// sample 3
	s = `abc <err>err-text</>
def <info>info text
</>`
	r = ReplaceTag(s)
	is.NotContains(r, "<")
	is.NotContains(r, ">")

	// sample 4
	s = "abc <err>err-text</> def <err>err-text</> "
	r = ReplaceTag(s)
	is.NotContains(r, "<")
	is.NotContains(r, ">")

	// sample 5
	s = "abc <err>err-text</> def <d>"
	r = ReplaceTag(s)
	is.NotContains(r, "<err>")
	is.Contains(r, "<d>")

	// sample 6
	s = "custom tag: <fg=yellow;bg=black;op=underscore;>hello, welcome</>"
	r = ReplaceTag(s)
	is.NotContains(r, "<")
	is.NotContains(r, ">")

	// empty
	s = Render()
	is.Equal("", s)
}

func TestTagParser_Parse_hex_rgb_c256(t *testing.T) {
	is := assert.New(t)
	p := NewTagParser()

	s := "custom tag: <fg=e7b2a1>hello, welcome</>"
	r := p.Parse(s)
	is.NotContains(r, "<")
	is.NotContains(r, ">")
	is.Equal("custom tag: \x1B[38;2;231;178;161mhello, welcome\x1B[0m", r)

	s = "custom tag: <fg=e7b2a1;bg=176;op=bold>hello, welcome</>"
	r = p.Parse(s)
	is.NotContains(r, "<")
	is.NotContains(r, ">")
	is.Equal("custom tag: \x1b[38;2;231;178;161;48;5;176;1mhello, welcome\x1b[0m", r)
}

func TestParseCodeFromAttr_basic(t *testing.T) {
	is := assert.New(t)

	s := ParseCodeFromAttr("=")
	is.Equal("", s)

	s = ParseCodeFromAttr("fg=lightRed;bg=lightRed;op=bold,blink")
	is.Equal("91;101;1;5", s)

	s = ParseCodeFromAttr("fg= lightRed;bg=lightRed;op=bold,")
	is.Equal("91;101;1", s)

	s = ParseCodeFromAttr("fg =lightRed;bg=lightRed;op=bold,blink")
	is.Equal("91;101;1;5", s)

	s = ParseCodeFromAttr("fg = lightRed;bg=lightRed;op=bold,blink")
	is.Equal("91;101;1;5", s)
}

func TestParseCodeFromAttr_c256(t *testing.T) {
	is := assert.New(t)

	s := ParseCodeFromAttr("fg=34")
	is.Equal("38;5;34", s)

	s = ParseCodeFromAttr("bg=56")
	is.Equal("48;5;56", s)

	s = ParseCodeFromAttr("fg=175; bg=56")
	is.Equal("38;5;175;48;5;56", s)
}

func TestParseCodeFromAttr_hex_rgb(t *testing.T) {
	is := assert.New(t)

	// --- hex code

	s := ParseCodeFromAttr("fg=e7b2a1")
	is.Equal("38;2;231;178;161", s)

	s = ParseCodeFromAttr("fg=e7b2a1;bg=c2c2c2")
	is.Equal("38;2;231;178;161;48;2;194;194;194", s)

	// --- rgb code

	s = ParseCodeFromAttr("fg=231,178,161")
	is.Equal("38;2;231;178;161", s)

	s = ParseCodeFromAttr("bg=231,178,161")
	is.Equal("48;2;231;178;161", s)
}

func TestPrint(t *testing.T) {
	is := assert.New(t)

	s := Sprint()
	is.Equal("", s)

	// force open color render for testing
	buf := forceOpenColorRender()
	defer resetColorRender()

	is.True(len(GetColorTags()) > 0)
	is.True(IsDefinedTag("info"))
	is.Equal("0;32", GetTagCode("info"))
	is.Equal("", GetTagCode("not-exist"))

	s = Sprint("<red>MSG</>")
	is.Equal("\x1b[0;31mMSG\x1b[0m", s)

	s = Sprint("<red>H</><green>I</>")
	is.Equal("\x1b[0;31mH\x1b[0m\x1b[0;32mI\x1b[0m", s)

	s = Sprintf("<red>%s</>", "MSG")
	is.Equal("\x1b[0;31mMSG\x1b[0m", s)

	// Print
	Print("<red>MSG</>")
	is.Equal("\x1b[0;31mMSG\x1b[0m", buf.String())
	buf.Reset()

	// Printf
	Printf("<red>%s</>", "MSG")
	is.Equal("\x1b[0;31mMSG\x1b[0m", buf.String())
	buf.Reset()

	// Println
	Println("<red>MSG</>")
	is.Equal("\x1b[0;31mMSG\x1b[0m\n", buf.String())
	buf.Reset()

	Println("<red>hello</>", "world")
	is.Equal("\x1b[0;31mhello\x1b[0m world\n", buf.String())
	buf.Reset()

	// Fprint
	Fprint(buf, "<red>MSG</>")
	is.Equal("\x1b[0;31mMSG\x1b[0m", buf.String())
	buf.Reset()

	// Fprintln
	Fprintln(buf, "<red>MSG</>")
	is.Equal("\x1b[0;31mMSG\x1b[0m\n", buf.String())
	buf.Reset()
	Fprintln(buf, "<red>hello</>", "world")
	is.Equal("\x1b[0;31mhello\x1b[0m world\n", buf.String())
	buf.Reset()

	// Fprintf
	Fprintf(buf, "<red>%s</>", "MSG")
	is.Equal("\x1b[0;31mMSG\x1b[0m", buf.String())
	buf.Reset()

	// Lprint
	logger := log.New(buf, "", 0)
	Lprint(logger, "<red>MSG</>\n")
	is.Equal("\x1b[0;31mMSG\x1b[0m\n", buf.String())
	buf.Reset()

	NotRenderTag()
	Fprintf(buf, "<red>%s</>", "MSG")
	is.Equal("<red>MSG</>", buf.String())
	buf.Reset()

	ResetOptions()
}

func TestWrapTag(t *testing.T) {
	at := assert.New(t)
	at.Equal("<info>text</>", WrapTag("text", "info"))
	at.Equal("", WrapTag("", "info"))
	at.Equal("text", WrapTag("text", ""))
}

func TestApplyTag(t *testing.T) {
	forceOpenColorRender()
	defer resetColorRender()
	at := assert.New(t)
	at.Equal("\x1b[0;32mMSG\x1b[0m", ApplyTag("info", "MSG"))
}

func TestClearTag(t *testing.T) {
	is := assert.New(t)
	is.Equal("text", ClearTag("text"))
	is.Equal("text", ClearTag("<err>text</>"))
	is.Equal("abc error def info text", ClearTag("abc <err>error</> def <info>info text</>"))

	str := `abc <err>err-text</>
def <info>info text
</>`
	ret := ClearTag(str)
	is.Contains(ret, "def info")
	is.NotContains(ret, "</>")

	str = "abc <err>text</> def<d>"
	ret = ClearTag(str)
	is.Equal("abc text def", ret)
	is.NotContains(ret, "<err>")
}

func TestTag_Print(t *testing.T) {
	buf := forceOpenColorRender()
	defer resetColorRender()
	is := assert.New(t)

	s := Tag("info").Sprint("msg")
	is.Equal("\x1b[0;32mmsg\x1b[0m", s)

	s = Tag("info").Sprintf("m%s", "sg")
	is.Equal("\x1b[0;32mmsg\x1b[0m", s)

	info := Tag("info")

	// Tag.Print
	info.Print("msg")
	is.Equal("\x1b[0;32mmsg\x1b[0m", buf.String())
	buf.Reset()

	// Tag.Println
	info.Println("msg")
	is.Equal("\x1b[0;32mmsg\x1b[0m\n", buf.String())
	buf.Reset()

	// Tag.Printf
	info.Printf("m%s", "sg")
	is.Equal("\x1b[0;32mmsg\x1b[0m", buf.String())
	buf.Reset()

	mga := Tag("mga")

	// Tag.Print
	mga.Print("msg")
	is.Equal("\x1b[0;35mmsg\x1b[0m", buf.String())
	buf.Reset()

	// Tag.Println
	mga.Println("msg", "more")
	is.Equal("\x1b[0;35mmsg more\x1b[0m\n", buf.String())
	buf.Reset()

	// Tag.Printf
	mga.Printf("m%s", "sg")
	is.Equal("\x1b[0;35mmsg\x1b[0m", buf.String())
	buf.Reset()
}
