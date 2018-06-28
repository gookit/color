package color

import (
	"fmt"
)

//Style a colored style
type Style struct {
	colors []Color // fg color, bg color, color options
}

// Apply
// usage:
// 	`(string, fg-color,bg-color, options...)`
//  color.Apply("text", color.FgGreen)
//  color.Apply("text", color.FgGreen, color.BgBlack, color.OpBold)
func Apply(str string, colors ...Color) string {
	return buildColoredText(
		buildColorCode(colors...),
		str,
	)
}

// New create a custom style
func New(colors ...Color) *Style {
	return &Style{colors}
}

// Colors
func (s Style) Colors() (colors []Color) {
	return s.colors
}

// Render render text
// usage:
//  color.New(color.FgGreen).Render("text")
//  color.New(color.FgGreen, color.BgBlack, color.OpBold).Render("text")
func (s Style) Render(args ...interface{}) string {
	return buildColoredText(
		buildColorCode(s.Colors()...),
		fmt.Sprint(args...),
	)
}

// Apply is alias of the 'Render'
func (s Style) Apply(args ...interface{}) string {
	return s.Render(args...)
}

// Print render and Print text
func (s Style) Print(args ...interface{}) (int, error) {
	return fmt.Print(s.Render(args...))
}

// Printf render and Print text
func (s Style) Printf(format string, args ...interface{}) (int, error) {
	str := fmt.Sprintf(format, args...)

	return fmt.Print(s.Render(str))
}

// Println render and Print text
func (s Style) Println(args ...interface{}) (int, error) {
	return fmt.Println(s.Render(args...))
}
