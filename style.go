package color

import (
	"fmt"
)

// Style a colored style
// can add: fg color, bg color, color options
// quick use:
// 	color.Style(color.FgGreen).
type Style []Color

// New create a custom style
func New(colors ...Color) Style {
	return Style(colors)
}

// Save save to styles map
func (s Style) Save(name string) {
	AddStyle(name, s)
}

// Render render text
// usage:
//  color.New(color.FgGreen).Render("text")
//  color.New(color.FgGreen, color.BgBlack, color.OpBold).Render("text")
func (s Style) Render(args ...interface{}) string {
	if isLikeInCmd {
		return fmt.Sprint(args...)
	}

	return buildColoredText(buildColorCode(s...), args...)
}

// Apply is alias of the 'Render'
func (s Style) Apply(args ...interface{}) string {
	return s.Render(args...)
}

// Print render and Print text
func (s Style) Print(args ...interface{}) (int, error) {
	if isLikeInCmd {
		return winPrint(fmt.Sprint(args...), s...)
	}

	return fmt.Print(s.Render(args...))
}

// Printf render and Print text
func (s Style) Printf(format string, args ...interface{}) (int, error) {
	str := fmt.Sprintf(format, args...)

	if isLikeInCmd {
		return winPrint(str, s...)
	}

	return fmt.Print(s.Render(str))
}

// Println render and Print text
func (s Style) Println(args ...interface{}) (int, error) {
	if isLikeInCmd {
		return winPrintln(fmt.Sprint(args...), s...)
	}

	return fmt.Println(s.Render(args...))
}
