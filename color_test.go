package color

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Example() {
	// simple usage
	Cyan.Printf("Simple to use %s\n", "color")

	// use like func
	red := FgRed.Render
	green := FgGreen.Render
	fmt.Printf("%s line %s library\n", red("Command"), green("color"))

	// custom color
	New(FgWhite, BgBlack).Println("custom color style")

	// can also:
	Style{FgCyan, OpBold}.Println("custom color style")

	// internal theme/style:
	Info.Prompt("message")
	Info.Println("message")
	Warn.Println("message")
	Error.Println("message")

	// use style tag
	Print("<suc>he</><comment>llo</>, <cyan>wel</><red>come</>\n")

	// set a style tag
	Tag("info").Println("info style text")

	// use info style tips
	Tips("info").Print("tips style text")

	// use info style blocked tips
	LiteTips("info").Print("blocked tips style text")
}

func TestColor_Render(t *testing.T) {
	at := assert.New(t)

	r := Bold.Render("text")
	at.Equal(r, "\x1b[1mtext\x1b[0m")
	// at.Equal(fmt.Sprintf("%q", r), "\x1b[1mtext\x1b[0m")
}

func TestRenderCodes(t *testing.T) {
	art := assert.New(t)
	art.Contains(RenderCodes("36;1", "Text"), "36;1")
}

func TestClearCode(t *testing.T) {
	art := assert.New(t)
	art.Equal("Text", ClearCode("\033[36;1mText\x1b[0m"))
	// 8bit
	art.Equal("Text", ClearCode("\x1b[38;5;242mText\x1b[0m"))
	// 24bit
	art.Equal("Text", ClearCode("\x1b[38;2;30;144;255mText\x1b[0m"))
	art.Equal("Text other", ClearCode("\033[36;1mText\x1b[0m other"))
}
