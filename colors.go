package color

// reset color
const Reset Color = 0

// Foreground colors.
const (
	// basic Foreground colors 30 - 37
	FgBlack   Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite

	FgDefault Color = 39

	// extra Foreground color 90 - 97
	FgDarkGray     Color = iota + 90
	FgLightRed
	FgLightGreen
	FgLightYellow
	FgLightBlue
	FgLightMagenta
	FgLightCyan
	FgWhiteEx
)

// Foreground colors map
var FgColors = map[string]Color{
	"black":   FgBlack,
	"red":     FgRed,
	"green":   FgGreen,
	"yellow":  FgYellow,
	"blue":    FgBlue,
	"magenta": FgMagenta,
	"cyan":    FgCyan,
	"white":   FgWhite,
	"default": FgDefault,
}

// Background colors.
const (
	// basic Background colors 40 - 47
	BgBlack   Color = iota + 40
	BgRed
	BgGreen
	BgYellow   // BgBrown like yellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
	BgDefault Color = 49

	// extra Background color 100 - 107
	BgDarkGray     Color = iota + 100
	BgLightRed
	BgLightGreen
	BgLightYellow
	BgLightBlue
	BgLightMagenta
	BgLightCyan
	BgWhiteEx
)

// Background colors map
var BgColors = map[string]Color{
	"black":   BgBlack,
	"red":     BgRed,
	"green":   BgGreen,
	"yellow":  BgYellow,
	"blue":    BgBlue,
	"magenta": BgMagenta,
	"cyan":    BgCyan,
	"white":   BgWhite,
	"default": BgDefault,
}

// color options
const (
	OpBold       = 1 // 加粗
	OpFuzzy      = 2 // 模糊(不是所有的终端仿真器都支持)
	OpItalic     = 3 // 斜体(不是所有的终端仿真器都支持)
	OpUnderscore = 4 // 下划线
	OpBlink      = 5 // 闪烁
	OpReverse    = 7 // 颠倒的 交换背景色与前景色
	OpConcealed  = 8 // 隐匿的
)

// color options map
var Options = map[string]Color{
	"bold":       OpBold,
	"fuzzy":      OpFuzzy,
	"italic":     OpItalic,
	"underscore": OpUnderscore,
	"blink":      OpBlink,
	"reverse":    OpReverse,
	"concealed":  OpConcealed,
}

// IsFgColor
func IsFgColor(name string) bool {
	if _, ok := FgColors[name]; ok {
		return true
	}

	return false
}

// IsBgColor
func IsBgColor(name string) bool {
	if _, ok := BgColors[name]; ok {
		return true
	}

	return false
}

// IsOption
func IsOption(name string) bool {
	if _, ok := Options[name]; ok {
		return true
	}

	return false
}
