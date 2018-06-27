package color

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	// Regex to match color tags
	// golang 不支持反向引用.  即不支持使用 \1 引用第一个匹配 ([a-z=;]+)
	// TagExpr = `<([a-z=;]+)>(.*?)<\/\1>`
	// 所以调整一下 统一使用 </> 来结束标签 e.g <info>some text</>
	// (?s:...) s - 匹配换行
	TagExpr = `<([a-z=;]+)>(?s:(.*?))<\/>`

	// Regex used for removing color tags
	// StripExpr = `<[\/]?[a-zA-Z=;]+>`
	// 随着上面的做一些调整
	StripExpr = `<[\/]?[a-zA-Z=;]*>`
)

//type Style string
type Style struct {
	Fg  Color   // fg color
	Bg  Color   // bg color
	Ops []Color // color options
}

// NewStyle create a custom style
func NewStyle(colors ...Color) *Style {
	//return Style(buildColorCode(colors...))
	switch len(colors) {
	case 0:
		return &Style{}
	case 1: // only fg
		return &Style{Fg: colors[0]}
	case 2: // only fg, bg
		return &Style{Fg: colors[0], Bg: colors[1]}
	default: // full fg, bg, opts...
		return &Style{Fg: colors[0], Bg: colors[1], Ops: colors[2:]}
	}
}

// Colors
func (stl Style) Colors() (colors []Color) {
	if stl.Fg > 0 {
		colors = append(colors, stl.Fg)
	}

	if stl.Bg > 0 {
		colors = append(colors, stl.Bg)
	}

	if len(stl.Ops) > 0 {
		colors = append(colors, stl.Ops...)
	}

	return
}

// Text render text
func (stl Style) Text(str string) string {
	return buildColoredText(
		buildColorCode(stl.Colors()...),
		str,
	)
}

// UseStyle
func UseStyle(name string, str string) string {
	return Render(WrapTag(str, name))
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

// Render return rendered string
func Render(str string) string {
	return ReplaceTag(str)
}

// ReplaceTag return rendered string
func ReplaceTag(str string) string {
	if !strings.Contains(str, "<") {
		return str
	}

	//reg := regexp.MustCompile(TagExpr)
	reg, err := regexp.Compile(TagExpr)

	if err != nil {
		return str
	}

	matches := reg.FindAllStringSubmatch(str, -1)
	// fmt.Printf("matches %v\n", matches)

	for _, item := range matches {
		// e.g "<tag>content</>"
		_, tag, content := item[0], item[1], item[2]
		code := GetStyleCode(tag)

		if len(code) > 0 {
			now := buildColoredText(code, content)
			old := WrapTag(content, tag)
			str = strings.Replace(str, old, now, 1)
		}
		// fmt.Printf("tag: %s, tag content:%s\n", tag, content)
	}

	return str
}

// IsStyle is style name
func IsStyle(name string) bool {
	if _, ok := Styles[name]; ok {
		return true
	}

	return false
}

// GetStyleCode get color code by style name
func GetStyleCode(name string) string {
	if code, ok := Styles[name]; ok {
		return code
	}

	return ""
}

// WrapTag wrap a tag for a string "<tag>content</>"
func WrapTag(str string, tag string) string {
	// return fmt.Sprintf("<%s>%s</%s>", tag, str, tag)
	return fmt.Sprintf("<%s>%s</>", tag, str)
}

// ClearTag clear all tag for a string
func ClearTag(str string) string {
	if !strings.Contains(str, "<") {
		return str
	}

	rgp, err := regexp.Compile(StripExpr)
	if err != nil {
		return str
	}

	return rgp.ReplaceAllString(str, "")
}
