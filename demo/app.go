package main

import (
	"github.com/gookit/color"
	"fmt"
)

func main() {
	colorUsage()
}

func colorUsage() {
	// simple usage
	color.FgCyan.Printf("Simple to use %s\n", "color")

	// custom color
	color.New(color.FgMagenta, color.BgBlack).Println("custom color style")
	// can also:
	color.Style{color.FgCyan, color.OpBold}.Println("custom color style")

	// use defined color tag
	color.Print("<suc>he</><comment>llo</>, <cyan>wel</><red>come</>\n")

	// use custom color tag
	color.Print("<fg=yellow;bg=black;op=underscore;>hello, welcome</>\n")

	// set a color tag
	color.Tag("info").Println("info style message")

	// tips
	color.Tips("info").Print("tips style message")
	color.Tips("warn").Print("tips style message")

	// lite tips
	color.LiteTips("info").Print("lite tips style message")
	color.LiteTips("warn").Print("lite tips style message")

	i := 0

	fmt.Print("\n- All Available color Tags: \n\n")

	for tag, _ := range color.GetColorTags() {
		i++
		color.Tag(tag).Print(tag)

		if i%5 == 0 {
			fmt.Print("\n")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}
