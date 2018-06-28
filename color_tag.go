package color

import (
	"strings"
	"regexp"
	"fmt"
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

// UseTagColor
func UseTagColor(name string, str string) string {
	return buildColoredText(GetStyleCode(name), str)
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

// IsDefinedTag is defined tag name
func IsDefinedTag(name string) bool {
	if _, ok := TagColors[name]; ok {
		return true
	}

	return false
}

// GetStyleCode get color code by tag name
func GetStyleCode(name string) string {
	if code, ok := TagColors[name]; ok {
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

// Some defined style tags, in the TagColors.
const (
	// basic
	TagRed     = "red"
	TagBlue    = "blue"
	TagCyan    = "cyan"
	TagBlack   = "black"
	TagGreen   = "green"
	TagBrown   = "brown"
	TagWhite   = "white"
	TagNormal  = "normal" // no color
	TagYellow  = "yellow"
	TagMagenta = "magenta"

	// alert
	Suc     = "suc" // same "green" and "bold"
	Success = "success"
	Info    = "info"    // same "green"
	Comment = "comment" // same "brown"
	Note    = "note"
	Notice  = "notice"
	Warn    = "warn"
	Warning = "warning"
	Danger  = "danger" // same "red"
	Err     = "err"
	Error   = "error"

	// option
	TagBold       = "bold"
	TagUnderscore = "underscore"
	TagReverse    = "reverse"
)

// Some internal defined style tags
// format is: "fg;bg;opt"
// usage: <tag>content text</>
var TagColors = map[string]string{
	// basic
	"red":     "0;31",
	"blue":    "0;34",
	"cyan":    "0;36",
	"black":   "0;30",
	"green":   "0;32",
	"brown":   "0;33",
	"white":   "1;37",
	"default": "39", // no color
	"normal":  "39", // no color
	"yellow":  "1;33",
	"magenta": "1;35",

	// alert
	"suc":     "1;32", // same "green" and "bold"
	"success": "1;32",
	"info":    "0;32", // same "green",
	"comment": "0;33", // same "brown"
	"note":    "36;1",
	"notice":  "36;4",
	"warn":    "0;30;43",
	"warning": "0;30;43",
	"danger":  "0;31", // same "red"
	"err":     "30;41",
	"error":   "30;41",

	// more
	"lightRed":      "1;31",
	"light_red":     "1;31",
	"lightGreen":    "1;32",
	"light_green":   "1;32",
	"lightBlue":     "1;34",
	"light_blue":    "1;34",
	"lightCyan":     "1;36",
	"light_cyan":    "1;36",
	"lightDray":     "37",
	"light_gray":    "37",
	"darkDray":      "90",
	"dark_gray":     "90",
	"lightYellow":   "93",
	"light_yellow":  "93",
	"lightMagenta":  "95",
	"light_magenta": "95",

	// extra
	"lightRedEx":     "91",
	"light_red_ex":   "91",
	"lightGreenEx":   "92",
	"light_green_ex": "92",
	"lightBlueEx":    "94",
	"light_blue_ex":  "94",
	"lightCyanEx":    "96",
	"light_cyan_ex":  "96",
	"whiteEx":        "97",
	"white_ex":       "97",

	// option
	"bold":       "1",
	"underscore": "4",
	"reverse":    "7",
}
