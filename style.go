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
	return buildColoredText(buildColorCode(s...), args...)
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

// style name
// type StyleName string

// Some defined style tags, in the BuiltinStyles/TagColors.
const (
	// alert tag, like bootstrap's alert
	Suc     = "suc" // same "green" and "bold"
	Success = "success"
	Info    = "info"    // same "green"
	Comment = "comment" // same "brown"
	Note    = "note"
	Notice  = "notice"
	Warn    = "warn"
	Warning = "warning"
	Primary = "primary"
	Danger  = "danger" // same "red"
	Err     = "err"
	Error   = "error"
)

// some built-in style list
var BuiltinStyles = map[string]Style{
	"info": {OpReset, FgGreen},
	"note": {OpBold, FgLightCyan},
	"error": {FgLightWhite, BgRed},
	"danger": {OpBold, FgRed},
	"notice": {OpBold, FgCyan},
	"success": {OpBold, FgGreen},
	"comment": {OpReset, FgYellow},
	"primary": {OpReset, FgBlue},
	"warning": {OpReset, FgBlack, BgLightYellow},
}

// some style name alias
var styleAliases = map[string]string{
	"err": "error",
	"suc": "success",
	"warn": "warning",
}

func AddStyle(name string, s Style) {
	BuiltinStyles[name] = s
}

// GetStyle get style by name
func GetStyle(name string) Style {
	if s, ok := BuiltinStyles[name]; ok {
		return s
	}

	if realName, ok := styleAliases[name]; ok {
		return BuiltinStyles[realName]
	}

	panic(fmt.Sprintf("style name '%s' is not exists", name))
}
